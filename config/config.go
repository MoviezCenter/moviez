package config

type AppConfig struct {
	DBConfig `mapstructure:",squash"`
}

type DBConfig struct {
	Host     string `mapstructure:"MOVIEZ_DB_HOST"`
	Database string `mapstructure:"MOVIEZ_DB_NAME"`
	Username string `mapstructure:"MOVIEZ_DB_USER"`
	Password string `mapstructure:"MOVIEZ_DB_PASSWORD"`
	Port     string `mapstructure:"MOVIEZ_DB_PORT"`
}

var AppConfigInstance AppConfig
