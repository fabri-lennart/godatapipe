package models

import (
	"time"
)

type Warehouse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LocationCity string `json:"location_city"`

	CapacityM3          float64 `json:"capacity_m3"`
	CurrentOccupancyPct float64 `json:"current_occupancy"`

	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
