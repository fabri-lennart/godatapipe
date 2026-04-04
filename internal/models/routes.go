package models

import "time"

type Route struct {
	ID string `json:"id"`

	OriginWarehouseID      string `json:"origin_warehouse_id"`
	DestinationWarehouseID string `json:"destination_warehouse_id"`

	OriginWarehouse      *Warehouse `json:"origin_warehouse,omitempty"`
	DestinationWarehouse *Warehouse `json:"destination_warehouse,omitempty"`

	DistanceKM      float64 `json:"distance_km"`
	AvgTransitHours float64 `json:"avg_transit_hours"`
	AvgCostBRL      float64 `json:"avg_cost_brl"`
	OnTimeRate      float64 `json:"on_time_rate"`

	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}
