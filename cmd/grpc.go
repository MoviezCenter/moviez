package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/MoviezCenter/moviez/config"
	"github.com/MoviezCenter/moviez/internal/controller"
	moviepb "github.com/MoviezCenter/pb-contracts-go/movie"
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

	grpcServer := grpc.NewServer()
	movieServiceServer := controller.NewMovieServiceServer()
	moviepb.RegisterMovieServiceServer(grpcServer, movieServiceServer)

	_, err := config.InitDB(config.AppConfigInstance.DBConfig)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

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
