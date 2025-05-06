package input

import (
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
)

type ProductCategoryService interface {
	CreateProductCategory(pc *dto.ProductCategoryDto) error

	GetProductCategoryById(id uint) (*models.ProductCategory, error)

	ListProductCategories(page int, pageSize int) ([]models.ProductCategory, error)
}
