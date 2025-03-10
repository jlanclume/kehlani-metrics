// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"
)

type MetricEvent struct {
	EventTimestamp time.Time
	ClientID       string
	TenantID       string
	MetricValue    float64
	MetricType     string
	MetricName     string
}

type MetricEventsHyper struct {
	EventTimestamp time.Time
	ClientID       string
	TenantID       string
	MetricValue    float64
	MetricType     string
	MetricName     string
}
