package config

import (
	"fmt"
	"os"
)

var dbConfig *databaseConfig

type databaseConfig struct {
	name     string
	host     string
	port     string
	user     string
	password string
}

func (dc databaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dc.host, dc.user, dc.password, dc.name, dc.port,
	)
}

func GetDatabaseConfig() *databaseConfig {
	if dbConfig == nil {
		dbConfig = &databaseConfig{
			name: os.Getenv("DB_NAME"),
			host: os.Getenv("DB_HOST"),
			port: os.Getenv("DB_PORT"),
			user: os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
		}
	}

	return dbConfig
}
