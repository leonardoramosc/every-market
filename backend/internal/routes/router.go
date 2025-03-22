package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _router *gin.Engine

func getRouter() *gin.Engine {
	if (_router == nil) {
		_router = gin.Default()
	}
	return _router
}

func InitRouter() {
	router := getRouter()
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "¡Hola desde Go Backend EveryMarket prueba 3"})
	})

	registerRoutes()

	log.Fatal(router.Run(":5000"))
}