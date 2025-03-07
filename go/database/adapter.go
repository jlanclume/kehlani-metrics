package database

import (
	"context"
	"kehlani_metrics/protobuf"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseAdapter struct {
	pool *pgxpool.Pool
}

// NewDatabaseAdapter creates a new adapter
func NewDatabaseAdapter(databaseUrl string) (*DatabaseAdapter, error) {
	config, err := pgxpool.ParseConfig(databaseUrl)

	if err != nil {
		log.Printf("Failed to parse config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Printf("Failed to instanciate a database adapter: %v", err)
		return nil, err
	}

	return &DatabaseAdapter{
		pool: pool,
	}, nil
}

// Free resources used by adapted
func (adapter *DatabaseAdapter) Close() {
	adapter.pool.Close()
}

// Persist a metric event
func (adapter *DatabaseAdapter) SaveMetricEvent(ctx context.Context, event *protobuf.MetricEvent) error {
	conn, err := adapter.pool.Acquire(ctx)

	if err != nil {
		log.Printf("Failed to acquire connection: %v", err)
		return err
	}

	defer conn.Release()

	queries := New(conn)
	_, err = queries.SaveMetricEvent(ctx, SaveMetricEventParams{
		EventTimestamp: time.Unix(event.EventTimestamp, 0),
		ClientID:       event.ClientId,
		TenantID:       event.TenantId,
		MetricValue:    event.MetricValue,
		MetricType:     event.MetricType.String(),
		MetricName:     event.MetricName,
	})

	if err != nil {
		log.Printf("Failed to execute SaveMetricEvent query: %v", err)
	}
	return err
}

// ListAllMetricEventFromDb fetches all metric events, unpaginated
func (adapter *DatabaseAdapter) ListAllMetricEventFromDb(ctx context.Context) ([]*protobuf.MetricEvent, error) {
	conn, err := adapter.pool.Acquire(ctx)
	if err != nil {
		log.Printf("Failed to acquire connection: %v", err)
		return nil, err
	}

	defer conn.Release()

	queries := New(conn)
	events, err := queries.ListAllMetricEvent(ctx)

	if err != nil {
		log.Printf("Failed to execute ListAllMetricEventFromDb query: %v", err)
		return nil, err
	}
	return convertDbMetricEventsHyperSlice(events), nil
}
