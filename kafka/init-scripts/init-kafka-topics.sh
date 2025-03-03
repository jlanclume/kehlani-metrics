#!/bin/bash

echo "Waiting for Kafka to be up and running"
kafka-topics.sh --bootstrap-server kafka:9092 --list

echo "Creating Kafka topic 'metric-event'"
kafka-topics.sh --bootstrap-server kafka:9092 --create --topic metric-event --partitions 3 --replication-factor 1
