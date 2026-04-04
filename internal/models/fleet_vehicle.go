package models

import "time"

type FleetVehicle struct {
	ID                 string `json:"id"`
	VehicleType        string `json:"vehicle_type"`
	Status             string `json:"status"`
	CurrentWarehouseID string `json:"current_warehouse_id"`

	CurrentWarehouse *Warehouse `json:"current_warehouse,omitempty"`

	LastServiceDate *time.Time `json:"last_service_date,omitempty"`
	Active          bool       `json:"active"`
}
