package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type WorkerConfig struct {
	KafkaBroker  string
	KafkaGroupID string
	DatabaseURL  string
}

// LoadWorkerConfig loads the configuration from environment variables
func LoadWorkerConfig() (*WorkerConfig, error) {
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

	config := &WorkerConfig{
		KafkaBroker:  getEnv("WORKER_KAFKA_BROKER", "kafka:9092"),
		KafkaGroupID: getEnv("WORKER_KAFKA_GROUP_ID", "db-writer-group"),
		DatabaseURL:  getEnv("WORKER_DATABASE_URL", ""),
	}

	// Additional checks for required environment variables
	if config.DatabaseURL == "" {
		return nil, errors.New("error: WORKER_DATABASE_URL is required but not set")
	}

	return config, nil
}
