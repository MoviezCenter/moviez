package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	pbRunTime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/MoviezCenter/moviez/config"
	"github.com/MoviezCenter/moviez/ent"
	"github.com/MoviezCenter/moviez/ent/user"
	moviepb "github.com/MoviezCenter/pb-contracts-go/movie"
)

var wait time.Duration = time.Second * 15

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start http gateway server",
	Run:   runHttpCmd,
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

func runHttpCmd(cmd *cobra.Command, args []string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	pbMux := pbRunTime.NewServeMux()
	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := moviepb.RegisterMovieServiceHandlerFromEndpoint(context.Background(), pbMux, ":8081", grpcOpts)
	if err != nil {
		panic(err)
	}
	httpMux := mux.NewRouter()
	httpMux.PathPrefix("/").Handler(pbMux)

	// health check
	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		str := "ok"
		w.Write([]byte(str))
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	httpHandler := cors.AllowAll().Handler(httpMux)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", config.AppConfigInstance.HTTPPort),
		Handler: httpHandler,
	}

	entClient, err := config.InitEntClient(config.AppConfigInstance.DBConfig)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}
	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := CreateGraph(context.Background(), entClient); err != nil {
		log.Fatalf("failed creating dummy data: %v", err)
	}

	if err := GetMovies(context.Background(), entClient); err != nil {
		log.Fatalf("failed get movies data: %v", err)
	}

	go func() {
		log.Printf("server is listing at port %s\n", config.AppConfigInstance.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error starting http server: %s", err.Error())
		}
	}()

	// Block till receiving the signal
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown returned error: %s", err.Error())
	}

	log.Println("server gracefully shutdown")
	os.Exit(0)
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	movie1, err := client.Movie.
		Create().
		SetTitle("Avengers").
		SetOverview("A team of remarkable people has to gather to save the earth").
		SetReleaseDate("2012-10-01").
		SetPosterPath("").
		Save(ctx)
	if err != nil {
		return err
	}

	movie2, err := client.Movie.
		Create().
		SetTitle("Avengers: Endgame").
		SetOverview("A team of remarkable people has to gather to save the earth").
		SetReleaseDate("2019-10-01").
		SetPosterPath("").
		Save(ctx)
	if err != nil {
		return err
	}

	review1, err := client.Review.
		Create().
		SetComment("Good movie").
		SetRating(4.5).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	review2, err := client.Review.
		Create().
		SetComment("Mediocre").
		SetRating(2.5).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	review3, err := client.Review.
		Create().
		SetComment("Great movie").
		SetRating(4.8).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	review4, err := client.Review.
		Create().
		SetComment("Ok movie").
		SetRating(3.5).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}

	movie1.Update().AddReviews(review1, review2).Save(ctx)
	movie2.Update().AddReviews(review3, review4).Save(ctx)

	_, err = client.User.
		Create().
		SetUsername("user01").
		SetEmail("user01@gmail.com").
		SetPassword("123456").
		AddReviews(review1, review3).
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = client.User.
		Create().
		SetUsername("user02").
		SetEmail("user02@gmail.com").
		SetPassword("123456").
		AddReviews(review2, review4).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetMovies(ctx context.Context, client *ent.Client) error {
	u1, err := client.User.Query().
		Where(
			user.HasReviews(),
			user.Username("user01"),
		).Only(ctx)
	if err != nil {
		return err
	}

	movies, err := u1.QueryReviews().QueryMovie().All(ctx)
	if err != nil {
		return err
	}

	log.Printf("movies user %s has reviewd", u1.Username)
	for _, m := range movies {
		log.Println("Title: ", m.Title)
	}

	return nil
}
