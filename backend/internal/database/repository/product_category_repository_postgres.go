package repository

import (
	"errors"

	"github.com/leonardoramosc/every-market/internal/database"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"gorm.io/gorm"
)

type productCategoryRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *productCategoryRepositoryPostgres) CreateProductCategory(pc *models.ProductCategory) error {
	result := repo.db.Create(pc)
	return result.Error
}

func (repo *productCategoryRepositoryPostgres) GetProductCategoryByName(name string) (*models.ProductCategory, error) {
	var pc models.ProductCategory
	result := repo.db.Where("name = ?", name).First(&pc)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &pc, result.Error
}

func (repo *productCategoryRepositoryPostgres) GetProductCategoryById(id int) (*models.ProductCategory, error) {
	var pc models.ProductCategory
	result := repo.db.First(&pc, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &pc, result.Error
}

func NewProductCategoryRepositoryPostgres() *productCategoryRepositoryPostgres {
	db := database.GetDatabaseClient()
	return &productCategoryRepositoryPostgres{db}
}
