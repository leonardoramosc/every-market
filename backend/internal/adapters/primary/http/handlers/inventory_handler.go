package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/core/ports/input"
	"github.com/leonardoramosc/every-market/internal/exceptions"
)

type InventoryHandler struct {
	inventoryService input.InventoryService
}

func NewInventoryHandler(service input.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService: service}
}

func (h *InventoryHandler) CreateInventory(ctx *gin.Context) {
	var request dto.CreateInventoryDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	i, err := h.inventoryService.CreateInventory(&request)

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