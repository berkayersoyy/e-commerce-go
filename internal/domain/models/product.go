package models

import "github.com/jinzhu/gorm"

const (
	ProductNotFound = "product not found"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required,min=5,max=45"`
	Price       float32 `json:"price" validate:"required"`
	CategoryId  uint    `json:"categoryId" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

//json incele
