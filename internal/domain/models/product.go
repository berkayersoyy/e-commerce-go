package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=5,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}
