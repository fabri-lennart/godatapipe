package usecase

import (
	"context"

	"github.com/fabri-lennart/godatapipe/internal/models"
	"github.com/fabri-lennart/godatapipe/internal/repository"
)

type WarehouseUseCase struct {
	repo repository.WarehouseRepository
}

func NewWarehouseUseCase(repo repository.WarehouseRepository) *WarehouseUseCase {
	return &WarehouseUseCase{repo: repo}
}

func (uc *WarehouseUseCase) GetAll(ctx context.Context) ([]models.Warehouse, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *WarehouseUseCase) GetByID(ctx context.Context, id string) (*models.Warehouse, error) {
	return uc.repo.GetByID(ctx, id)
}
