package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBSource string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file", err)
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE is not set")
	}

	return &Config{DBSource: dbSource}, nil
}
