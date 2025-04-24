package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/dto"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/internal/services"
)

func CreateProductHandler(ctx *gin.Context) {
	service := services.NewProductService()
	categoryService := services.NewProductCategoryService()
	var request dto.CreateProductDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := categoryService.GetProductCategoryById(request.CategoryID)

	if err != nil {
		message := "unable to create product"
		log.Printf("\n%v. Error trying to validate category Reason: %v\n", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	if category == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf("category with id=%v does not exist", request.CategoryID)})
		return
	}

	newProduct, err := service.CreateProduct(&request)

	if err != nil {
		message := "unable to create product"
		log.Printf("\n%v. Reason: %v\n", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	response := responses.ProductResponse{
		Id:          uint(newProduct.ID),
		Name:        newProduct.Name,
		Description: newProduct.Description,
		ImageURL:    newProduct.ImageURL,
		Price:       newProduct.Price,
		CategoryID:  newProduct.ProductCategoryID,
	}
	ctx.JSON(http.StatusOK, response)
}
