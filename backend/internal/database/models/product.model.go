package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	ImageURL    string

	ProductCategoryID int
	ProductCategory   ProductCategory

	Inventory Inventory

	ProductImages []ProductImage // one-to-many
}

type ProductImage struct {
	gorm.Model
	ProductID uint
	URL       string
}
