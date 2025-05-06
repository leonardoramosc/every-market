package main

import (
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/handlers"
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/routes"
	database "github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/repositories"
	"github.com/leonardoramosc/every-market/internal/core/services"
)

func main() {
	db := database.Connect()
	database.RunMigrations()

	inventoryRepo := repositories.NewGormInventoryRepositoryPostgres(db)
	productRepo := repositories.NewGormProductRepositoryPostgres(db)
	productCategoryRepo := repositories.NewGormProductCategoryRepositoryPostgres(db)

	productService := services.NewProductService(productRepo)
	inventoryService := services.NewInventoryService(inventoryRepo, productService)
	productCategoryService := services.NewProductCategoryService(productCategoryRepo)

	productHandler := handlers.NewProductHandler(productService, productCategoryService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)
	ProductCategoryHandler := handlers.NewProductCategoryHandler(productCategoryService)

	handlersToRegister := &handlers.Handlers{
		ProductHandler: productHandler, 
		InventoryHandler: inventoryHandler, 
		ProductCategoryHandler: ProductCategoryHandler,
	}

	routes.InitRouter(handlersToRegister)
}
