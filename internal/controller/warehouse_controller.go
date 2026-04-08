package controller

import (
	"net/http"

	"github.com/fabri-lennart/godatapipe/internal/usecase"
	"github.com/gin-gonic/gin"
)

type WarehouseController struct {
	useCase *usecase.WarehouseUseCase
}

func NewWarehouseController(uc *usecase.WarehouseUseCase) *WarehouseController {
	return &WarehouseController{useCase: uc}
}

func (ctrl *WarehouseController) GetAll(c *gin.Context) {
	warehouses, err := ctrl.useCase.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   warehouses,
	})
}

func (ctrl *WarehouseController) GetByID(c *gin.Context) {
	id := c.Param("id")

	warehouse, err := ctrl.useCase.GetByID(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   warehouse,
	})
}
