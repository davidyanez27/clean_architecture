package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID            int
	Code          string
	Name          string
	ProductTypeID int
	UnitMeasureID int
	State         bool
	ProductType   ProductType
	UnitMeasure   UnitMeasure
}
