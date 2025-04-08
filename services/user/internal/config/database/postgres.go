package database

import (
	"log"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  bool
	TimeZone string
}

func NewPostgresConfig() *PostgresConfig {
	envPort := os.Getenv("POSTGRES_PORT")
	port, err := strconv.Atoi(envPort)

	if err != nil {
		log.Fatalf("Invalid database port number: %s", envPort)
	}
	if port <= 0 || port > 65535 {
		log.Fatalf("Database port number out of range: %d", port)
	}

	return &PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE") == "true",
		TimeZone: os.Getenv("POSTGRES_TIMEZONE"),
	}
}
