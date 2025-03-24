package dto

type CreateInventoryDto struct {
	Stock int `json:"stock" binding:"required,numeric,gte=0"`
	ProductID int `json:"productId" binding:"required"`
}
