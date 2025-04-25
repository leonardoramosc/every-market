package domain

type Inventory struct {
	Stock     int
	ProductID int
}


type ProductCategory struct {
  Name         string
}

type ProductImage struct {
	ProductID uint
	URL       string
}

type Product struct {
	Name        string
	Description string
	Price       float64
	ImageURL    string

	CategoryID uint
	Category   ProductCategory

	Inventory Inventory

	Images []ProductImage
}


