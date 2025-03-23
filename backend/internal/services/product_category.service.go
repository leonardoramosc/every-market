package services

import (
	"strings"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"github.com/leonardoramosc/every-market/internal/database/repository"
	"github.com/leonardoramosc/every-market/internal/dto"
	"github.com/leonardoramosc/every-market/internal/exceptions"
)

var productCategoryService *_productCategoryService

type _productCategoryService struct {
	repo repository.ProductCategoryRepository
}

func (pcs *_productCategoryService) CreateProductCategory(pc *dto.ProductCategoryDto) error {
	name := strings.ToLower(pc.Name)
	model := &models.ProductCategory{Name: name}
	existingProductCategory, err := pcs.repo.GetProductCategoryByName(name)
	if err != nil {
		return err
	}
	if existingProductCategory != nil {
		return exceptions.ErrProductCategoryExists
	}
	return pcs.repo.CreateProductCategory(model)
}

func (pcs *_productCategoryService) GetProductCategoryById(id int) (*models.ProductCategory, error) {
	return pcs.repo.GetProductCategoryById(id)
}

func NewProductCategoryService() *_productCategoryService {
	if productCategoryService == nil {
		repo := repository.NewProductCategoryRepositoryPostgres()
		productCategoryService = &_productCategoryService{repo}
	}
	return productCategoryService
}
