package repositories

import (
	"errors"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
	"gorm.io/gorm"
)

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

func NewGormInventoryRepositoryPostgres(db *gorm.DB) *inventoryRepositoryPostgres {
	return &inventoryRepositoryPostgres{db}
}
