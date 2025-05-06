package services

import (
	"strings"

	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
	"github.com/leonardoramosc/every-market/internal/core/ports/output"
	"github.com/leonardoramosc/every-market/internal/exceptions"
)

var productCategoryService *_productCategoryService

type _productCategoryService struct {
	repo output.ProductCategoryRepository
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

func (pcs *_productCategoryService) GetProductCategoryById(id uint) (*models.ProductCategory, error) {
	return pcs.repo.GetProductCategoryById(id)
}

func (pcs *_productCategoryService) ListProductCategories(page int, pageSize int) ([]models.ProductCategory, error) {
	return pcs.repo.ListProductCategories(page, pageSize)
}

func NewProductCategoryService(repo output.ProductCategoryRepository) *_productCategoryService {
	if productCategoryService == nil {
		productCategoryService = &_productCategoryService{repo}
	}
	return productCategoryService
}
