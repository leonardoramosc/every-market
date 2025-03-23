package services

import (
	"github.com/leonardoramosc/every-market/internal/database/models"
	"github.com/leonardoramosc/every-market/internal/database/repository"
	"github.com/leonardoramosc/every-market/internal/dto"
)

var productService *_productService

type _productService struct {
	repo repository.ProductRepository
}

func (ps *_productService) CreateProduct(product *dto.CreateProductDto) (*models.Product, error) {
	pm := models.Product{
		Name:        product.Name,
		Description: product.Description,
		ImageURL:    product.ImageURL,
		Price:       product.Price,
		ProductCategoryID: product.CategoryID,
	}
	return ps.repo.CreateProduct(&pm)
}

func NewProductService() *_productService {
	if productService == nil {
		repo := repository.NewProductRepositoryPostgres()
		productService = &_productService{repo}
	}
	return productService
}
