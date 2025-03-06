package main

import (
	"context"
	"kehlani_metrics/database"
	"kehlani_metrics/internal/config"
	"kehlani_metrics/messaging"
	"kehlani_metrics/protobuf"
	"log"
	"net"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"google.golang.org/protobuf/proto"
)

type server struct {
	protobuf.UnimplementedMetricEventServiceServer
	kafkaProducer *messaging.KafkaProducer
	dbAdapter     *database.DatabaseAdapter
}

func (s *server) SaveMetricEvent(ctx context.Context, req *protobuf.MetricEventRequest) (*protobuf.MetricEventResponse, error) {
	// Extract metric event from the request
	event, err := proto.Marshal(req.GetEvent())
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return nil, err
	}

	// Build Kafka message
	msg := &kafka.Message{
		Topic: "metric-event",
		Value: event,
	}

	// Publsh message
	if err := s.kafkaProducer.ProduceMessage(ctx, msg); err != nil {
		log.Printf("Failed to produce event: %v", err)
		return nil, err
	}

	response := &protobuf.MetricEventResponse{
		Status: "OK",
	}

	return response, nil
}

func main() {
	// Load configuration
	config, err := config.LoadApiConfig()

	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Create a new Kafka producer
	kafkaProducer, err := messaging.NewKafkaProducer(config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	// Create a new DB adapter
	dbAdapter, err := database.NewDatabaseAdapter(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to create a Database adapter: %v", err)
	}
	defer dbAdapter.Close()

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Enable reflection for testing
	reflection.Register(grpcServer)

	// Register the metric service with the server
	protobuf.RegisterMetricEventServiceServer(grpcServer, &server{
		kafkaProducer: kafkaProducer,
		dbAdapter:     dbAdapter,
	})

	log.Println("Starting server on port 50051...")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
