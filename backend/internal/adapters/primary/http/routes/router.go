package routes

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/handlers"
)

var _router *gin.Engine

func getRouter() *gin.Engine {
	if _router == nil {
		_router = gin.Default()
	}
	return _router
}

func InitRouter(handlersToRegister *handlers.Handlers) {
	router := getRouter()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // TODO: Cambiar esto por la URL del frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "¡Hola desde Go Backend EveryMarket"})
	})

	registerRoutes(handlersToRegister)

	log.Fatal(router.Run(":5000"))
}
