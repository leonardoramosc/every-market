package exceptions

import "errors"

var ErrProductCategoryExists = errors.New("product category already exists")
var ErrProductNotExistForInventory = errors.New("product does not exist")
var ErrInventoryAlreadyExistForProduct = errors.New("inventory already exist")
