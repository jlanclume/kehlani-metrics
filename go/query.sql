-- name: SaveMetricEvent :one
INSERT INTO metric_events(event_timestamp, client_id, tenant_id, metric_value, metric_type, metric_name)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
    )
RETURNING *;

-- name: GetMetricEvent :one
SELECT * from metric_events_hyper 
where client_id = $1 
limit 1;

-- name: ListAllMetricEvent :many
SELECT * from metric_events_hyper;
