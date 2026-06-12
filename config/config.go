package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func LoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Provide a default PostgreSQL fallback
		dbURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  port,
	}
}
