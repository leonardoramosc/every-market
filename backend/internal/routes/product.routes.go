package routes

import (
	"github.com/leonardoramosc/every-market/internal/handlers"
)

func registerProductRoutes() {
	router := getRouter()
	productCategoryRoutes := router.Group("/api/products")
	productCategoryRoutes.GET("/", handlers.ListProductsHandler)
	productCategoryRoutes.GET("/:id", handlers.GetProductByIdHandler)
	productCategoryRoutes.POST("/", handlers.CreateProductHandler)
	productCategoryRoutes.GET("/category/:category", handlers.ListProductsByCategoryHandler)
}
