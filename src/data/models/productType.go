package models

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	ID      int
	Code    string
	Name    string
	Product []Product
}
