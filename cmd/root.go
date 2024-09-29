package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/TranTheTuan/go-template/config"
)

const (
	DBHost = "db_host"
	DBPort = "db_port"
	DBUser = "db_user"
	DBPass = "db_password"
	DBName = "db_name"
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
	initConfiguration()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config.AppConfigInstance); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func initConfiguration() {
	rootCmd.PersistentFlags().String(DBHost, "localhost", "Database host")
	rootCmd.PersistentFlags().String(DBPort, "5432", "Database port")
	rootCmd.PersistentFlags().String(DBUser, "postgres", "Database user")
	rootCmd.PersistentFlags().String(DBPass, "postgres", "Database password")
	rootCmd.PersistentFlags().String(DBName, "postgres", "Database name")

	viper.BindPFlag(DBHost, rootCmd.PersistentFlags().Lookup(DBHost))
	viper.BindPFlag(DBPort, rootCmd.PersistentFlags().Lookup(DBPort))
	viper.BindPFlag(DBUser, rootCmd.PersistentFlags().Lookup(DBUser))
	viper.BindPFlag(DBPass, rootCmd.PersistentFlags().Lookup(DBPass))
	viper.BindPFlag(DBName, rootCmd.PersistentFlags().Lookup(DBName))
}
