package config

import (
	"fmt"
	"os"

	"github.com/mcuadros/go-defaults"
)

type DBConfig struct {
	Host     string `default:"0.0.0.0"`
	Port     string `default:"3036"`
	User     string `default:"root"`
	DBName   string `default:"cuongpo"`
	Password string `default:"cuongpo"`
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DBNAME"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	defaults.SetDefaults(&dbConfig)
	return &dbConfig
}

func BuildDBConn() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
