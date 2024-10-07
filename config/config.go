package config

type AppConfig struct {
	HTTPPort string `mapstructure:"PORT"`
	DBConfig `mapstructure:",squash"`
}

var AppConfigInstance AppConfig
