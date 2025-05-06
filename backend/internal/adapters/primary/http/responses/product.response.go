package responses

import "time"

type ProductResponse struct {
	Id          uint       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"imageURL"`
	CategoryID  uint       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`

	Images []string `json:"images"`
	Stock int `json:"stock"` 
}
