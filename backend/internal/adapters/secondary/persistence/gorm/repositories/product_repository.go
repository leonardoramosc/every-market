package repositories

import (
	"errors"
	"log"

	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/gormutil"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
	"gorm.io/gorm"
)

type productRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *productRepositoryPostgres) CreateProduct(product *models.Product, images []string) (*models.Product, error) {
	result := repo.db.Create(&product)
	_, err := repo.CreateImages(product.ID, images)

	if err != nil {
		log.Printf("unable to create images for product %v\n", product.ID)
	}

	return product, result.Error
}

func (repo *productRepositoryPostgres) CreateImages(productID uint, images []string) ([]models.ProductImage, error) {
	var productImages []models.ProductImage

	for _, imageURL := range images {
		productImage := models.ProductImage{ProductID: productID, URL: imageURL}
		productImages = append(productImages, productImage)
	}

	result := repo.db.Create(&productImages)
	return productImages, result.Error
}

func (repo *productRepositoryPostgres) ListProducts(page int, pageSize int) (*[]models.Product, error) {
	var products []models.Product
	result := repo.db.Scopes(gormutil.Paginate(page, pageSize)).Preload("Inventory").Find(&products)

	return &products, result.Error
}

func (repo *productRepositoryPostgres) ListProductsByCategory(category string, page int, pageSize int) (*[]models.Product, error) {
	var products []models.Product
	result := repo.db.
		Joins("JOIN product_categories ON product_categories.id = products.product_category_id").
		Where("product_categories.name = ?", category).
		Scopes(gormutil.Paginate(page, pageSize)).
		Find(&products)

	return &products, result.Error
}

func (repo *productRepositoryPostgres) GetProductById(id int) (*models.Product, error) {
	var p models.Product
	result := repo.db.Preload("ProductImages").First(&p, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &p, result.Error
}

func NewGormProductRepositoryPostgres(db *gorm.DB) *productRepositoryPostgres {
	return &productRepositoryPostgres{db}
}
