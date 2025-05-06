package httputil

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(ctx *gin.Context) (page int, pageSize int) {
	var err error

	page, err = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, err = strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	if err != nil {
		log.Printf("\nunable to extract page & page_size from query string. Reason: %v\n", err.Error())
		return 1, 10
	}

	return page, pageSize
}
