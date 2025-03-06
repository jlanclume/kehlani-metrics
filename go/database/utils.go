package database

import "kehlani_metrics/protobuf"

// Convert a single db metric event to proto metric event
func convertDbMetricEventsHyper(dbEvent MetricEventsHyper) *protobuf.MetricEvent {
	return &protobuf.MetricEvent{
		EventTimestamp: dbEvent.EventTimestamp.UnixMilli(),
		ClientId:       dbEvent.ClientID,
		TenantId:       dbEvent.TenantID,
		MetricValue:    dbEvent.MetricValue,
		MetricType:     protobuf.MetricType(protobuf.MetricType_value[dbEvent.MetricType]),
		MetricName:     dbEvent.MetricName,
	}
}

// Convert a slice of db metric events to a proto metric event
func convertDbMetricEventsHyperSlice(dbEvents []MetricEventsHyper) []*protobuf.MetricEvent {
	var events []*protobuf.MetricEvent
	for _, dbEvent := range dbEvents {
		events = append(events, convertDbMetricEventsHyper(dbEvent))
	}
	return events
}
