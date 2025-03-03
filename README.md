Kehlani Metrics
===============

Kehlani Metrics is a data analytics project designed to collect, process, and analyze real-time metrics. This project leverages various technologies such as Kafka, TimescaleDB, and Docker to create an efficient data pipeline.

Features
--------

- **Kafka (Kraft)**: Used as the messaging broker to stream data in real-time.
- **TimescaleDB**: A PostgreSQL extension for time-series data, allowing efficient storage and querying of metrics.
- **Docker Compose**: Used to set up and manage the different services in the project.

Getting Started
---------------

To start the components of this project, run the following command:
```sh
docker compose up
```

This will initialize all necessary services in Docker, including Kafka, TimescaleDB, and others.
