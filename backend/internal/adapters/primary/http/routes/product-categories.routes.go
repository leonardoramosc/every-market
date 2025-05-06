package routes

import "github.com/leonardoramosc/every-market/internal/adapters/primary/http/handlers"

func registerProductCategoriesRoutes(handler *handlers.ProductCategoryHandler) {
	router := getRouter()
	productCategoryRoutes := router.Group("/api/product-categories")
	productCategoryRoutes.POST("/", handler.CreateProductCategory)
	productCategoryRoutes.GET("/", handler.ListProductCategories)
}
