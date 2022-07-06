package models

import "github.com/jinzhu/gorm"

const (
	CategoryNotFound = "category not found"
)

type Category struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=2,max=45"`
}
