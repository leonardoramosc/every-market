package routes

import "github.com/leonardoramosc/every-market/internal/adapters/primary/http/handlers"

func registerRoutes(handlersToRegister *handlers.Handlers) {
	registerProductCategoriesRoutes(handlersToRegister.ProductCategoryHandler)
	registerProductRoutes(handlersToRegister.ProductHandler)
	registerInventoryRoutes(handlersToRegister.InventoryHandler)
}