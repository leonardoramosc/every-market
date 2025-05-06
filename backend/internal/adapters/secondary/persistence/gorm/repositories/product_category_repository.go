package repositories

import (
	"errors"

	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/gormutil"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
	"gorm.io/gorm"
)

type productCategoryRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *productCategoryRepositoryPostgres) CreateProductCategory(pc *models.ProductCategory) error {
	result := repo.db.Create(pc)
	return result.Error
}

func (repo *productCategoryRepositoryPostgres) GetProductCategoryByName(name string) (*models.ProductCategory, error) {
	var pc models.ProductCategory
	result := repo.db.Where("name = ?", name).First(&pc)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &pc, result.Error
}

func (repo *productCategoryRepositoryPostgres) GetProductCategoryById(id uint) (*models.ProductCategory, error) {
	var pc models.ProductCategory
	result := repo.db.First(&pc, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &pc, result.Error
}

func (repo *productCategoryRepositoryPostgres) ListProductCategories(page int, pageSize int) ([]models.ProductCategory, error) {
	var categories []models.ProductCategory
	result := repo.db.Scopes(gormutil.Paginate(page, pageSize)).Find(&categories)

	return categories, result.Error
}

func NewGormProductCategoryRepositoryPostgres(db *gorm.DB) *productCategoryRepositoryPostgres {
	return &productCategoryRepositoryPostgres{db}
}
