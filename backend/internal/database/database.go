package database

import (
	"log"

	"github.com/leonardoramosc/every-market/config"
	"github.com/leonardoramosc/every-market/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() *gorm.DB {
	dbConfig := config.GetDatabaseConfig()
	dsn := dbConfig.GetDSN()

	var err error

	log.Println("Conneting to database...")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	log.Println("Connected successfully to database!")

	return db
}

func GetDatabaseClient() *gorm.DB {
	return db
}

func RunMigrations() {
	err := db.AutoMigrate(models.ProductCategory{}, models.Product{}, models.Inventory{}, models.ProductImage{})

	if err != nil {
		log.Fatal("Unable to run migrations", err.Error())
	}
}
