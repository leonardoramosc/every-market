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

	i, errors := service.CreateInventory(&request)

	if len(errors) > 0 {
		statusCode, formattedErrors := formatErrors(errors)
		ctx.JSON(statusCode, formattedErrors)
		return
	}

	ctx.JSON(http.StatusCreated, i)
}

func formatErrors(errs []error) (int, gin.H) {
	var formattedErrors []string
	statusCode := http.StatusUnprocessableEntity
	for _, err := range errs {
		if errors.Is(err, exceptions.ErrProductNotExistForInventory) {
			formattedErrors = append(formattedErrors, "product does not exist")
			continue
		}
		if errors.Is(err, exceptions.ErrInventoryAlreadyExistForProduct) {
			formattedErrors = append(formattedErrors, "inventory already exist for this product")
			continue
		}
		message := "unable to create inventory"
		formattedErrors = append(formattedErrors, message)
		statusCode = http.StatusInternalServerError
		log.Printf("\n%v. Reason: %v\n", message, err.Error())
	}
	return statusCode, gin.H{"errors": formattedErrors}
}
