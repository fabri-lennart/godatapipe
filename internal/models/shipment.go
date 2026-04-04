package models

import "time"

type Shipment struct {
	ID          string     `json:"id"`
	Carrer      *string    `json:"carrer,omitempty"`
	Weight      float64    `json:"weight"`
	ShippedAt   *time.Time `json:"shipped_at,omitempty"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	SLAMet      bool       `json:"slam_met"`
}
