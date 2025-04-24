package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/internal/services"
	"github.com/leonardoramosc/every-market/internal/utils"
)

func ListProductsHandler(ctx *gin.Context) {
	service := services.NewProductService()

	page, pageSize := utils.GetPaginationParams(ctx)

	products, err := service.ListProducts(page, pageSize)

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
