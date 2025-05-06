package responses

import "time"

type ProductCategoryResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
