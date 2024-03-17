package config

import (
	"fmt"
	"os"

	"github.com/mcuadros/go-defaults"
)

type DBConfig struct {
	Host     string `default:"0.0.0.0"`
	Port     string `default:"3036"`
	User     string `default:"cuongpo"`
	DBName   string `default:"cuongpo"`
	Password string `default:"cuongpo998"`
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_DBNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	defaults.SetDefaults(&dbConfig)
	return &dbConfig
}

func BuildDSN() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
