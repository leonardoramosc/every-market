package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/core/ports/input"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/pkg/httputil"
)

type ProductIDUri struct {
	ProductID string `uri:"id" binding:"required,numeric"`
}

type CategoryUri struct {
	Category string `uri:"category" binding:"required"`
}

type ProductHandler struct {
	productService input.ProductService
	productCategoryService input.ProductCategoryService
}

func NewProductHandler(service input.ProductService, productCategoryService input.ProductCategoryService) *ProductHandler {
	return &ProductHandler{productService: service, productCategoryService: productCategoryService}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var request dto.CreateProductDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.productCategoryService.GetProductCategoryById(request.CategoryID)

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

	newProduct, err := h.productService.CreateProduct(&request)

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

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	var uriParams ProductIDUri
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.Atoi(uriParams.ProductID)

	product, err := h.productService.GetProductById(id)

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

func (h *ProductHandler) ListProductsByCategory(ctx *gin.Context) {
	var uriParams CategoryUri
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	page, pageSize := httputil.GetPaginationParams(ctx)

	products, err := h.productService.ListProductsByCategory(uriParams.Category, page, pageSize)

	if err != nil {
		message := "unable to get products"
		log.Printf("\n%v. Reason: %v\n", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	res := []responses.ProductResponse{}

	for _, product := range *products {
		res = append(
			res,
			responses.ProductResponse{
				Id:          uint(product.ID),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				ImageURL:    product.ImageURL,
				CategoryID:  product.ProductCategoryID,
				CreatedAt:   product.CreatedAt,
			},
		)
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *ProductHandler) ListProducts(ctx *gin.Context) {
	page, pageSize := httputil.GetPaginationParams(ctx)

	products, err := h.productService.ListProducts(page, pageSize)

	if err != nil {
		message := "unable to get products"
		log.Printf("\n%v. Reason: %v\n", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	res := []responses.ProductResponse{}

	for _, product := range *products {
		res = append(
			res,
			responses.ProductResponse{
				Id:          uint(product.ID),
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				ImageURL:    product.ImageURL,
				CategoryID:  product.ProductCategoryID,
				CreatedAt:   product.CreatedAt,
			},
		)
	}

	ctx.JSON(http.StatusOK, res)
}