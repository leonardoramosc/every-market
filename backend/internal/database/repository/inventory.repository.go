package repository

import (
	"errors"

	"github.com/leonardoramosc/every-market/internal/database"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	CreateInventory(inventory *models.Inventory) (*models.Inventory, error)
	GetInventoryByProduct(productID int) (*models.Inventory, error)
}

type inventoryRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *inventoryRepositoryPostgres) CreateInventory(inventory *models.Inventory) (*models.Inventory, error) {
	result := repo.db.Create(inventory)
	return inventory, result.Error
}

func (repo *inventoryRepositoryPostgres) GetInventoryByProduct(productID int) (*models.Inventory, error) {
	var inventory models.Inventory
	err := repo.db.Where("product_id = ?", productID).First(&inventory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &inventory, nil
}

func NewInventoryRepositoryPostgres() *inventoryRepositoryPostgres {
	db := database.GetDatabaseClient()
	return &inventoryRepositoryPostgres{db}
}
