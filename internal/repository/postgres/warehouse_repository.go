package postgres

import (
	"context"
	"database/sql"

	"github.com/fabri-lennart/godatapipe/internal/models"
	"github.com/fabri-lennart/godatapipe/internal/repository"
)

type warehouseRepo struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) repository.WarehouseRepository {
	return &warehouseRepo{db: db}
}

func (r *warehouseRepo) Create(ctx context.Context, w *models.Warehouse) error {
	query := `INSERT INTO warehouses (id, name, location_city, capacity_m3, active, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, w.ID, w.Name, w.LocationCity, w.CapacityM3, w.Active, w.CreatedAt, w.UpdatedAt)
	return err
}

func (r *warehouseRepo) GetByID(ctx context.Context, id string) (*models.Warehouse, error) {
	query := `SELECT id, name, location_city, capacity_m3, current_occupancy_pct, active, created_at, updated_at
              FROM warehouses WHERE id = $1`

	var w models.Warehouse
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&w.ID, &w.Name, &w.LocationCity, &w.CapacityM3, &w.CurrentOccupancyPct, &w.Active, &w.CreatedAt, &w.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *warehouseRepo) GetAll(ctx context.Context) ([]models.Warehouse, error) {
	query := `SELECT id, name, location_city, capacity_m3, current_occupancy_pct, active, created_at, updated_at
              FROM warehouses`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var w models.Warehouse
		err := rows.Scan(
			&w.ID, &w.Name, &w.LocationCity, &w.CapacityM3,
			&w.CurrentOccupancyPct, &w.Active, &w.CreatedAt, &w.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}
	return warehouses, nil
}
