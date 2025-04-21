package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/internal/services"
	"github.com/leonardoramosc/every-market/internal/utils"
)

type CategoryUri struct {
	Category string `uri:"category" binding:"required"`
}

func ListProductsByCategoryHandler(ctx *gin.Context) {
	var uriParams CategoryUri
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	service := services.NewProductService()

	page, pageSize := utils.GetPaginationParams(ctx)

	products, err := service.ListProductsByCategory(uriParams.Category, page, pageSize)

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
				Id:          int(product.ID),
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
