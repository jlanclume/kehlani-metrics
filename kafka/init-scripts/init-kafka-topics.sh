#!/bin/bash

echo "Creating Kafka topic 'metric-event', if not exists"

# Using a single partition in dev mode to maintain a number of partitions lesser than the number of workers
kafka-topics.sh --bootstrap-server kafka:9092 --create --topic metric-event --partitions 1 --replication-factor 1 --if-not-exists
