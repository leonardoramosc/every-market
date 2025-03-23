package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/dto"
	"github.com/leonardoramosc/every-market/internal/exceptions"
	"github.com/leonardoramosc/every-market/internal/services"
)

func CreateProductCategoryHandler(ctx *gin.Context) {
	service := services.NewProductCategoryService()
	var createProductCategoryRequest dto.ProductCategoryDto

	if err := ctx.ShouldBindJSON(&createProductCategoryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateProductCategory(&createProductCategoryRequest)

	if err != nil {
		if errors.Is(err, exceptions.ErrProductCategoryExists) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "category already exist"})
			return
		}
		message := "unable to create product category"
		log.Printf("%v. Reason: %v", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"name": createProductCategoryRequest.Name})
}
