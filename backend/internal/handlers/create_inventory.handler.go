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

func CreateInventoryHandler(ctx *gin.Context) {
	service := services.NewInventoryService()
	var request dto.CreateInventoryDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	i, err := service.CreateInventory(&request)

	if err != nil {
		if errors.Is(err, exceptions.ErrProductNotExistForInventory) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "product does not exist"})
			return
		}
		if errors.Is(err, exceptions.ErrInventoryAlreadyExistForProduct) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "inventory already exist for this product"})
			return
		}
		message := "unable to create inventory"
		log.Printf("\n%v. Reason: %v\n", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	ctx.JSON(http.StatusCreated, i)
}
