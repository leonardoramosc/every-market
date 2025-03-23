package repository

import (
	"github.com/leonardoramosc/every-market/internal/database/models"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
}
