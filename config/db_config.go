package config

import (
	"fmt"

	"github.com/MoviezCenter/moviez/ent"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DriverName = "postgres"

type DBConfig struct {
	Host     string `mapstructure:"MOVIEZ_DB_HOST"`
	Database string `mapstructure:"MOVIEZ_DB_NAME"`
	Username string `mapstructure:"MOVIEZ_DB_USER"`
	Password string `mapstructure:"MOVIEZ_DB_PASSWORD"`
	Port     string `mapstructure:"MOVIEZ_DB_PORT"`
}

func InitDB(dbConfig DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)
	db, err := sqlx.Open(DriverName, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitEntClient(dbConfig DBConfig) (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return client, nil
}
