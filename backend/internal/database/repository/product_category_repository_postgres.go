package repository

import (
	"log"

	"github.com/leonardoramosc/every-market/internal/database"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"gorm.io/gorm"
)

type productCategoryRepositoryPostgres struct {
	db *gorm.DB
}

func (repo *productCategoryRepositoryPostgres) CreateProductCategory(pc *models.ProductCategory) error {
	if repo.db == nil {
		log.Println("+++++++++++++++++++++++++++++++++")
		log.Println("DATABASE IS NIL")
		log.Println("+++++++++++++++++++++++++++++++++")
	}
	result := repo.db.Create(&pc)
	return result.Error
}

func NewProductCategoryRepositoryPostgres() *productCategoryRepositoryPostgres {
	db := database.GetDatabaseClient()
	return &productCategoryRepositoryPostgres{db}
}
