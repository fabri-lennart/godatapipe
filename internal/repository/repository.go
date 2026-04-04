package repository

import (
	"context"

	"github.com/fabri-lennart/godatapipe/internal/models"
)

type WarehouseRepository interface {
	Create(ctx context.Context, w *models.Warehouse) error
	GetByID(ctx context.Context, id string) (*models.Warehouse, error)
	GetAll(ctx context.Context) ([]models.Warehouse, error)
}

type ShipmentRepository interface {
	Create(ctx context.Context, s *models.Shipment) error
	GetByID(ctx context.Context, id string) (*models.Shipment, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	GetByWarehouse(ctx context.Context, warehouseID string) ([]models.Shipment, error)
}
