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

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/MoviezCenter/moviez/config"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "run http server",
	Long:  "run http server",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		r := mux.NewRouter()
		r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			str := "hello world"
			w.Write([]byte(str))
			w.WriteHeader(http.StatusOK)
		}).Methods(http.MethodGet)

		srv := http.Server{
			Addr:    fmt.Sprintf(":%s", config.AppConfigInstance.HTTPPort),
			Handler: r,
		}

		go func() {
			log.Printf("server is listing at port %s\n", config.AppConfigInstance.HTTPPort)
			if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("server returned error: %s", err.Error())
			}
		}()

		<-c
		log.Println("got interruption signal")
		if err := srv.Shutdown(context.TODO()); err != nil {
			log.Fatalf("server shutdown returned error: %s", err.Error())
		}

		log.Println("server gracefully shutdown")
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
