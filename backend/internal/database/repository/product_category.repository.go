package repository

import (
	"github.com/leonardoramosc/every-market/internal/database/models"
)

type ProductCategoryRepository interface {
	CreateProductCategory(pc *models.ProductCategory) error
}
