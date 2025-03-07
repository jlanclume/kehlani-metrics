package messaging

import (
	"context"
	"log"

	"kehlani_metrics/internal/config"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

// KafkaProducer holds the Kafka producer instance.
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer initializes the Kafka producer and connects to the broker.
func NewKafkaProducer(config *config.ApiConfig) (*KafkaProducer, error) {
	// Create a new Kafka writer (producer) with the broker and topic
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{config.KafkaBroker},
		CompressionCodec: &compress.SnappyCodec,
		BatchSize:        10,
		BatchTimeout:     100 * 1000,
	})

	kp := &KafkaProducer{
		writer: writer,
	}

	return kp, nil
}

// ProduceMessage sends a message to the configured Kafka topic.
func (kp *KafkaProducer) ProduceMessage(ctx context.Context, msg *kafka.Message) error {

	// Write the message to Kafka
	err := kp.writer.WriteMessages(ctx, *msg)
	if err != nil {
		log.Printf("failed to produce message: %v", err)

		return err
	}

	log.Printf("Message successfully produced to topic %s", msg.Topic)
	return nil
}

// Close closes the Kafka writer (producer) to release resources.
func (kp *KafkaProducer) Close() {
	if err := kp.writer.Close(); err != nil {
		log.Printf("Error closing Kafka writer: %v", err)
	}
}
