package routes

import "github.com/leonardoramosc/every-market/internal/handlers"

func registerInventoryRoutes() {
	router := getRouter()
	inventoryRoutes := router.Group("/api/inventory")
	inventoryRoutes.POST("/", handlers.CreateInventoryHandler)
}