package config

type AppConfig struct {
	DBConfig
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Database string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Port     string `mapstructure:"DB_PORT"`
}

var AppConfigInstance DBConfig
