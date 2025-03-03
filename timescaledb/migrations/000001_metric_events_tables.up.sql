BEGIN;

-- Create a single for all tenants
CREATE TABLE metric_events (
    event_timestamp TIMESTAMPTZ NOT NULL,
    client_id TEXT NOT NULL,
    tenant_id TEXT NOT NULL,
    metric_value DOUBLE PRECISION NOT NULL,
    metric_type TEXT NOT NULL,
    metric_name TEXT NOT NULL,
    PRIMARY KEY (event_timestamp, tenant_id, metric_name)
);


-- Create hypertable counterpart
CREATE TABLE metric_events_hyper (
    event_timestamp TIMESTAMPTZ NOT NULL,
    client_id TEXT NOT NULL,
    tenant_id TEXT NOT NULL,
    metric_value DOUBLE PRECISION NOT NULL,
    metric_type TEXT NOT NULL,
    metric_name TEXT NOT NULL,
    PRIMARY KEY (event_timestamp, tenant_id, metric_name)
);

-- Convert it into a hypertable (partitioning by event_timestamp)
SELECT create_hypertable('metric_events_hyper', 'event_timestamp');

COMMIT;