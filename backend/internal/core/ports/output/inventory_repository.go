package output

import "github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"

type InventoryRepository interface {
	CreateInventory(inventory *models.Inventory) (*models.Inventory, error)
	GetInventoryByProduct(productID int) (*models.Inventory, error)
}
