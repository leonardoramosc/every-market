package services

import (
	"log"
	"sync"

	"github.com/leonardoramosc/every-market/internal/database/models"
	"github.com/leonardoramosc/every-market/internal/database/repository"
	"github.com/leonardoramosc/every-market/internal/dto"
	"github.com/leonardoramosc/every-market/internal/exceptions"
)

var inventoryService *_inventoryService

type _inventoryService struct {
	repo           repository.InventoryRepository
	productService *_productService
}

func (inventoryService *_inventoryService) CreateInventory(inventory *dto.CreateInventoryDto) (*models.Inventory, []error) {
	i := models.Inventory{
		ProductID: inventory.ProductID,
		Stock:     inventory.Stock,
	}

	errCh := inventoryService.validateInventoryForProduct(
		inventoryService.validateInventory(i.ProductID),
		inventoryService.validateProduct(i.ProductID),
	)

	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	newInventory, err := inventoryService.repo.CreateInventory(&i)

	if err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	return newInventory, nil
}

func (inventoryService *_inventoryService) validateInventoryForProduct(errorChannels ...<-chan error) <-chan error {
	outputCh := make(chan error)
	var wg sync.WaitGroup

	for _, errCh := range errorChannels {
		wg.Add(1)
		go func(ch <-chan error) {
			defer wg.Done()
			for err := range ch {
				outputCh <- err
			}
		}(errCh)
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}

func (inventoryService *_inventoryService) validateProduct(productId int) <-chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)
		product, productErr := inventoryService.productService.GetProductById(productId)
		if productErr != nil {
			log.Printf("\nunable to get product with id=%v\n", productId)
			ch <- productErr
		}

		if product == nil {
			ch <- exceptions.ErrProductNotExistForInventory
		}
	}()

	return ch
}

func (inventoryService *_inventoryService) validateInventory(productId int) <-chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)
		existingInventory, inventoryErr := inventoryService.repo.GetInventoryByProduct(productId)
		if inventoryErr != nil {
			log.Println("unable to validate if inventory already exist")
			ch <- inventoryErr
		}

		if existingInventory != nil {
			ch <- exceptions.ErrInventoryAlreadyExistForProduct
		}
	}()

	return ch
}

func NewInventoryService() *_inventoryService {
	if inventoryService == nil {
		repo := repository.NewInventoryRepositoryPostgres()
		productService := NewProductService()
		inventoryService = &_inventoryService{repo, productService}
	}
	return inventoryService
}
