package routes

import (
	"github.com/leonardoramosc/every-market/internal/handlers"
)

func registerProductCategoriesRoutes() {
	router := getRouter()
	productCategoryRoutes := router.Group("/api/product-categories")
	productCategoryRoutes.POST("/", handlers.CreateProductCategoryHandler)
}
