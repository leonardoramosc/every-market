package routes

import "github.com/leonardoramosc/every-market/internal/adapters/primary/http/handlers"

func registerProductRoutes(handler *handlers.ProductHandler) {
	router := getRouter()
	productCategoryRoutes := router.Group("/api/products")
	productCategoryRoutes.GET("/", handler.ListProducts)
	productCategoryRoutes.GET("/:id", handler.GetProductById)
	productCategoryRoutes.POST("/", handler.CreateProduct)
	productCategoryRoutes.GET("/category/:category", handler.ListProductsByCategory)
}
