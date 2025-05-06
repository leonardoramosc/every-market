package input

import (
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
)

type InventoryService interface {
	CreateInventory(inventory *dto.CreateInventoryDto) (*models.Inventory, error)
}