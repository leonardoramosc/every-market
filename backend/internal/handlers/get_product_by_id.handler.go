package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/internal/services"
)

type ProductIDUri struct {
	ProductID string `uri:"id" binding:"required,numeric"`
}

func GetProductByIdHandler(ctx *gin.Context) {
	var uriParams ProductIDUri
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewProductService()

	id, _ := strconv.Atoi(uriParams.ProductID)

	product, err := service.GetProductById(id)

	if err != nil {
		message := "unable to get product"
		log.Printf("%v. Reason: %v", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	if product == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product does not exist"})
	}

	var images []string

	for _, image := range product.ProductImages {
		images = append(images, image.URL)
	}

	response := responses.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageURL:    product.ImageURL,
		Price:       product.Price,
		CategoryID:  product.ProductCategoryID,
		Images: images,
	}
	ctx.JSON(http.StatusOK, response)
} 