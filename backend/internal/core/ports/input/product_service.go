package input

import (
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
)

type ProductService interface {
	CreateProduct(product *dto.CreateProductDto) (*models.Product, error)

	ListProducts(page int, pageSize int) (*[]models.Product, error)

	ListProductsByCategory(category string, page int, pageSize int) (*[]models.Product, error)

	GetProductById(id int) (*models.Product, error)
}
