package config

type AppConfig struct {
	DBConfig `mapstructure:",squash"`
}

type DBConfig struct {
	Host     string `mapstructure:"db-host"`
	Database string `mapstructure:"db-name"`
	Username string `mapstructure:"db-user"`
	Password string `mapstructure:"db-password"`
	Port     string `mapstructure:"db-port"`
}

var AppConfigInstance AppConfig
