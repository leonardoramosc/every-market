package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Stock     int
	ProductID int `gorm:"index"`
}
