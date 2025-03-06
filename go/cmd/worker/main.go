package main

import (
	"context"
	"log"

	"kehlani_metrics/database"
	"kehlani_metrics/internal/config"
	"kehlani_metrics/messaging"
)

func main() {
	// Load config
	cfg, err := config.LoadWorkerConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	dbAdapter, err := database.NewDatabaseAdapter(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error while creating DB Adapter: %v", err)
	}
	defer dbAdapter.Close()

	// Initialize the Kafka consumer
	consumer := messaging.NewKafkaConsumer(cfg.KafkaBroker, cfg.KafkaGroupID, dbAdapter)

	// Start consuming messages
	defer func() {
		err = consumer.Stop()

		if err != nil {
			log.Fatalf("Error while stopping consumer: %v", err)
		}
	}()

	ctx := context.Background()
	consumer.ConsumeMessages(ctx)
}
