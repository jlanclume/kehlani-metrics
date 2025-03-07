package messaging

import (
	"context"
	"kehlani_metrics/database"
	"kehlani_metrics/protobuf"
	"log"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type KafkaConsumer struct {
	reader    *kafka.Reader
	dbAdapter *database.DatabaseAdapter
}

// Free all resources used by kafka consumer
func (kc *KafkaConsumer) Stop() error {
	return kc.reader.Close()
}

// NewKafkaConsumer creates a new Kafka consumer.
func NewKafkaConsumer(broker string, groupID string, dbAdapter *database.DatabaseAdapter) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{broker},
		Topic:          "metric-event",
		GroupID:        groupID,
		CommitInterval: 1,
	})

	return &KafkaConsumer{
		reader:    reader,
		dbAdapter: dbAdapter,
	}
}

// ConsumeMessages consumes messages from Kafka.
func (kc *KafkaConsumer) ConsumeMessages(ctx context.Context) {
	log.Println("Consuming messages from Kafka topic...")

	for {
		msg, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Deserialize the Protobuf message
		var metricEvent protobuf.MetricEvent
		if err := proto.Unmarshal(msg.Value, &metricEvent); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		// Print the deserialized message
		log.Printf("Received message: tenantId=%s, clientId=%s, Timestamp=%d, Metric Name=%s, Metric Value=%f, Metric Type=%s",
			metricEvent.TenantId,
			metricEvent.ClientId,
			metricEvent.EventTimestamp,
			metricEvent.MetricName,
			metricEvent.MetricValue,
			metricEvent.MetricType)

		// Persist event
		if err := kc.dbAdapter.SaveMetricEvent(ctx, &metricEvent); err != nil {
			log.Printf("Failed to save metric event: %v", err)
		}
	}
}
