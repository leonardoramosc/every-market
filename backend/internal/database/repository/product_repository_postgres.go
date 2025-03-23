package repository

import (
	"github.com/leonardoramosc/every-market/internal/database"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"gorm.io/gorm"
)

type productRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *productRepositoryPostgres) CreateProduct(product *models.Product) (*models.Product, error) {
	result := repo.db.Create(product)
	return product, result.Error
}

func NewProductRepositoryPostgres() *productRepositoryPostgres {
	db := database.GetDatabaseClient()
	return &productRepositoryPostgres{db}
}
