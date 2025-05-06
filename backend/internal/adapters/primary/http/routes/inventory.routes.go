package routes

import (
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/handlers"
)

func registerInventoryRoutes(inventoryHandler *handlers.InventoryHandler) {
	router := getRouter()
	inventoryRoutes := router.Group("/api/inventory")
	inventoryRoutes.POST("/", inventoryHandler.CreateInventory)
}