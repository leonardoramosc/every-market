package responses

import "time"

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageURL"`
	CategoryID  int     `json:"categoryId"`
	CreatedAt time.Time `json:"createdAt"`
}
