package dto

type CreateProductDto struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price float64 `json:"price" binding:"required,numeric,gt=0"`
	ImageURL string `json:"imageURL" binding:"required,url"`
	CategoryID int `json:"categoryID" binding:"required"`
	Images []string `json:"images" binding:"required,min=3,dive,required"`
}