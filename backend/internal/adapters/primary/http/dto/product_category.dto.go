package dto

type ProductCategoryDto struct {
	Name string `json:"name" binding:"required"`
}