package output

import "github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"

type ProductCategoryRepository interface {
	CreateProductCategory(pc *models.ProductCategory) error
	GetProductCategoryByName(name string) (*models.ProductCategory, error)
	GetProductCategoryById(id uint) (*models.ProductCategory, error)
	ListProductCategories(page int, pageSize int) ([]models.ProductCategory, error)
}
