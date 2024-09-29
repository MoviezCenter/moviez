package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
  Use: "http",
  Short: "run http server",
  Long: "run http server",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Run http server")
  },
}

func init() {
  rootCmd.AddCommand(httpCmd)
}
