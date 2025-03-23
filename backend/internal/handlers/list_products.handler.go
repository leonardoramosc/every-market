package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/response"
)

func ListProductsHandler(ctx *gin.Context) {
	response := []response.ProductResponse{}
	ctx.JSON(http.StatusOK, response)
}
