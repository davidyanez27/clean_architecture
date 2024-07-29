package models

import "gorm.io/gorm"

type UnitMeasure struct {
	gorm.Model
	ID      int
	Code    string
	Name    string
	Product []Product
}
