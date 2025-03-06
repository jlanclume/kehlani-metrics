package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	KafkaBroker string
	DatabaseURL string
}

// LoadApiConfig loads the configuration from environment variables
func LoadApiConfig() (*ApiConfig, error) {
	// Optionally load .env file for local development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	getEnv := func(key string, defaultValue string) string {
		if value, exists := os.LookupEnv(key); exists {
			return value
		}
		return defaultValue
	}

	config := &ApiConfig{
		KafkaBroker: getEnv("API_KAFKA_BROKER", "kafka:9092"),
		DatabaseURL: getEnv("API_DATABASE_URL", ""),
	}

	// Additional checks for required environment variables
	if config.DatabaseURL == "" {
		return nil, errors.New("error: API_DATABASE_URL is required but not set")
	}

	return config, nil
}
