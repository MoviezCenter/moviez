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
	httpMux.Handle("/", pbMux)

	// health check
	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		str := "ok"
		w.Write([]byte(str))
		w.WriteHeader(http.StatusOK)
	})

	httpHandler := cors.AllowAll().Handler(httpMux)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", config.AppConfigInstance.HTTPPort),
		Handler: httpHandler,
	}

	_, err = config.InitDB(config.AppConfigInstance.DBConfig)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
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
