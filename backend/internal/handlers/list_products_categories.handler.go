package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoramosc/every-market/internal/responses"
	"github.com/leonardoramosc/every-market/internal/services"
	"github.com/leonardoramosc/every-market/internal/utils"
)

func ListProductCategoriesHandler(ctx *gin.Context) {
	fmt.Println("++++++++++++++ LLEGO AL HANDLER ++++++++++++++++")
	service := services.NewProductCategoryService()

	page, pageSize := utils.GetPaginationParams(ctx)

	categories, err := service.ListProductCategories(page, pageSize)

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
