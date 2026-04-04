package models

import "time"

type DeliveryEvent struct {
	ID         string `json:"id"`
	ShipmentID string `json:"shipment_id"`

	EventType           string    `json:"event_type"`
	Location            string    `json:"location"`
	OccurrenceTimestamp time.Time `json:"occurrence_timestamp"`
	Description         string    `json:"description"`

	Shipment *Shipment `json:"shipment,omitempty"`
}
