package services

import (
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
	"github.com/leonardoramosc/every-market/internal/core/ports/output"
)

var productService *_productService

type _productService struct {
	repo output.ProductRepository
}

func (ps *_productService) CreateProduct(product *dto.CreateProductDto) (*models.Product, error) {
	pm := models.Product{
		Name:              product.Name,
		Description:       product.Description,
		ImageURL:          product.ImageURL,
		Price:             product.Price,
		ProductCategoryID: product.CategoryID,
	}
	return ps.repo.CreateProduct(&pm, product.Images)
}

func (ps *_productService) ListProducts(page int, pageSize int) (*[]models.Product, error) {
	return ps.repo.ListProducts(page, pageSize)
}

func (ps *_productService) ListProductsByCategory(category string, page int, pageSize int) (*[]models.Product, error) {
	return ps.repo.ListProductsByCategory(category, page, pageSize)
}

func (ps *_productService) GetProductById(id int) (*models.Product, error) {
	return ps.repo.GetProductById(id)
}

func NewProductService(repo output.ProductRepository) *_productService {
	if productService == nil {
		productService = &_productService{repo}
	}
	return productService
}
