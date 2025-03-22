package services

import (
	"strings"

	"github.com/leonardoramosc/every-market/internal/database/models"
	"github.com/leonardoramosc/every-market/internal/database/repository"
	"github.com/leonardoramosc/every-market/internal/dto"
)

var productCategoryService *_productCategoryService

type IProductCategoryService interface {
	CreateProductCategory(pc *dto.ProductCategoryDto) error
}

type _productCategoryService struct {
	repo repository.ProductCategoryRepository
}

func (pcs *_productCategoryService) CreateProductCategory(pc *dto.ProductCategoryDto) error {
	name := strings.ToLower(pc.Name)
	model := &models.ProductCategory{Name: name}
	return pcs.repo.CreateProductCategory(model)
}

func GetProductCategoryService() *_productCategoryService {
	if productCategoryService == nil {
		repo := repository.NewProductCategoryRepositoryPostgres()
		productCategoryService = &_productCategoryService{repo}
	}
	return productCategoryService
}
