package config

type AppConfig struct {
	HTTPPort string `mapstructure:"MOVIEZ_HTTP_PORT"`
	DBConfig `mapstructure:",squash"`
}

var AppConfigInstance AppConfig
