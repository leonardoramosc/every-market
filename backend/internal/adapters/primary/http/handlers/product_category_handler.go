package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/adapters/primary/http/dto"
	"github.com/leonardoramosc/every-market/internal/core/ports/input"
	"github.com/leonardoramosc/every-market/internal/exceptions"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/pkg/httputil"
)

type ProductCategoryHandler struct {
	productCategoryService input.ProductCategoryService
}

func NewProductCategoryHandler(productCategoryService input.ProductCategoryService) *ProductCategoryHandler {
	return &ProductCategoryHandler{productCategoryService: productCategoryService}
}

func (h *ProductCategoryHandler) CreateProductCategory(ctx *gin.Context) {
	var createProductCategoryRequest dto.ProductCategoryDto

	if err := ctx.ShouldBindJSON(&createProductCategoryRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.productCategoryService.CreateProductCategory(&createProductCategoryRequest)

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

func (h * ProductCategoryHandler) ListProductCategories(ctx *gin.Context) {
	page, pageSize := httputil.GetPaginationParams(ctx)

	categories, err := h.productCategoryService.ListProductCategories(page, pageSize)

	if err != nil {
		message := "unable to list product categories"
		log.Printf("%v. Reason: %v", message, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	res := []responses.ProductCategoryResponse{}

	for _, category := range categories {
		res = append(
			res,
			responses.ProductCategoryResponse{
				Id:        int(category.ID),
				Name:      category.Name,
				CreatedAt: category.CreatedAt,
			},
		)
	}

	ctx.JSON(http.StatusOK, res)
}
