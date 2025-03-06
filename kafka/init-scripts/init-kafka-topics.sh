#!/bin/bash

echo "Creating Kafka topic 'metric-event', if not exists"
kafka-topics.sh --bootstrap-server kafka:9092 --create --topic metric-event --partitions 3 --replication-factor 1 --if-not-exists
