package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/MoviezCenter/moviez/config"
)

const (
	DBHost     = "MOVIEZ_DB_HOST"
	DBName     = "MOVIEZ_DB_NAME"
	DBPort     = "MOVIEZ_DB_PORT"
	DBPassword = "MOVIEZ_DB_PASSWORD"
	DBUser     = "MOVIEZ_DB_USER"
	DBHttpPort = "MOVIEZ_HTTP_PORT"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root command",
		Long:  "Root command",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	viper.BindEnv(DBHost)
	viper.BindEnv(DBUser)
	viper.BindEnv(DBPassword)
	viper.BindEnv(DBName)
	viper.BindEnv(DBPort)
	viper.BindEnv(DBHttpPort, "PORT")
}

func initConfig() {
	viper.SetEnvPrefix("moviez")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config.AppConfigInstance); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
