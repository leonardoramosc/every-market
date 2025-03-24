package repository

import (
	"errors"

	"github.com/leonardoramosc/every-market/internal/database"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	ListProducts(page int, pageSize int) (*[]models.Product, error)
	GetProductById(id int) (*models.Product, error)
}

type productRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *productRepositoryPostgres) CreateProduct(product *models.Product) (*models.Product, error) {
	result := repo.db.Create(product)
	return product, result.Error
}

func (repo *productRepositoryPostgres) ListProducts(page int, pageSize int) (*[]models.Product, error) {
	var products []models.Product
	result := repo.db.Scopes(database.Paginate(page, pageSize)).Preload("Inventory").Find(&products)

	return &products, result.Error
}

func (repo *productRepositoryPostgres) GetProductById(id int) (*models.Product, error) {
	var p models.Product
	result := repo.db.First(&p, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &p, result.Error
}

func NewProductRepositoryPostgres() *productRepositoryPostgres {
	db := database.GetDatabaseClient()
	return &productRepositoryPostgres{db}
}
