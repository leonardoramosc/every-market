package models

type ProductCategory struct {
	ID           uint `json:"id"`
  Name         string `json:"name" gorm:"uniqueIndex"`
}