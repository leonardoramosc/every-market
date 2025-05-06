package services

import (
	"log"
	"sync"

	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/adapters/secondary/persistence/gorm/models"
	"github.com/leonardoramosc/every-market/internal/core/ports/input"
	"github.com/leonardoramosc/every-market/internal/core/ports/output"
	"github.com/leonardoramosc/every-market/internal/exceptions"
)

var inventoryService *_inventoryService

type _inventoryService struct {
	repo           output.InventoryRepository
	productService input.ProductService
}

func (ps *_inventoryService) CreateInventory(inventory *dto.CreateInventoryDto) (*models.Inventory, error) {
	i := models.Inventory{
		ProductID: inventory.ProductID,
		Stock:     inventory.Stock,
	}
	var product *models.Product
	var existingInventory *models.Inventory
	var productErr, inventoryErr error

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		product, productErr = ps.productService.GetProductById(i.ProductID)
	}()

	go func() {
		defer wg.Done()
		existingInventory, inventoryErr = ps.repo.GetInventoryByProduct(i.ProductID)
	}()

	wg.Wait()

	if productErr != nil {
		log.Printf("\nunable to get product with id=%v\n", i.ProductID)
		return nil, productErr
	}

	if inventoryErr != nil {
		log.Println("unable to validate if inventory already exist")
		return nil, inventoryErr
	}

	if existingInventory != nil {
		return nil, exceptions.ErrInventoryAlreadyExistForProduct
	}

	if product == nil {
		return nil, exceptions.ErrProductNotExistForInventory
	}

	return ps.repo.CreateInventory(&i)
}

func NewInventoryService(repo output.InventoryRepository, productService input.ProductService) *_inventoryService {
	if inventoryService == nil {
		inventoryService = &_inventoryService{repo, productService}
	}
	return inventoryService
}
