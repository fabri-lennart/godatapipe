package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fabri-lennart/godatapipe/internal/database"
	"github.com/fabri-lennart/godatapipe/internal/repository/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Database Connection
	db, err := database.NewPostgresConnection(
		"localhost", "5432", "postgres", "postgress", "postgres",
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 2. Initialize Repository
	warehouseRepo := postgres.NewWarehouseRepository(db)

	// 3. Setup Router
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Test Endpoint: Get all warehouses directly from Repo
	r.GET("/warehouses", func(c *gin.Context) {
		warehouses, err := warehouseRepo.GetAll(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, warehouses)
	})

	r.Run(":8080")
}
