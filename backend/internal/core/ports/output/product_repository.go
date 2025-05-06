package output

import "github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"

type ProductRepository interface {
	CreateProduct(product *models.Product, images []string) (*models.Product, error)
	ListProducts(page int, pageSize int) (*[]models.Product, error)
	ListProductsByCategory(category string, page int, pageSize int) (*[]models.Product, error)
	GetProductById(id int) (*models.Product, error)
}
