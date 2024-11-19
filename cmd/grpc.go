package cmd

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/MoviezCenter/moviez/config"
	"github.com/MoviezCenter/moviez/internal/controller"
	"github.com/MoviezCenter/moviez/internal/repository"
	"github.com/MoviezCenter/moviez/internal/service"
	moviepb "github.com/MoviezCenter/pb-contracts-go/movie"
	reviewpb "github.com/MoviezCenter/pb-contracts-go/review"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run:   runGrpcCmd,
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

func runGrpcCmd(cmd *cobra.Command, args []string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	entClient, err := config.InitEntClient(config.AppConfigInstance.DBConfig)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}
	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// repository
	movieRepo := repository.NewMovieRepo(entClient)
	reviewRepo := repository.NewReviewRepository(entClient)

	// service
	movieService := service.NewMovieService(movieRepo)
	reviewService := service.NewReviewService(reviewRepo)

	// controller
	movieServiceServer := controller.NewMovieServiceServer(movieService)
	reviewServiceServer := controller.NewReviewServiceServer(reviewService)

	// register grpc server
	grpcServer := grpc.NewServer()
	moviepb.RegisterMovieServiceServer(grpcServer, movieServiceServer)
	reviewpb.RegisterReviewServiceServer(grpcServer, reviewServiceServer)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("error listing to tcp port 8081: %s", err.Error())
	}

	go func() {
		log.Printf("grpc server is serving at port %s\n", "8081")
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("error starting grpc server: %s", err.Error())
		}
	}()

	// Block till receiving the signal
	<-c
	grpcServer.GracefulStop()

	log.Println("server gracefully shutdown")
	os.Exit(0)
}
