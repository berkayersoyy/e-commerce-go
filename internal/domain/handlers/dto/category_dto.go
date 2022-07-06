package dto

import (
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/jinzhu/gorm"
	"time"
)

//CreateCategoryDto Create category dto
type CreateCategoryDto struct {
	Name string `json:"name" validate:"required,min=5,max=45"`
}

//UpdateCategoryDto Update category dto
type UpdateCategoryDto struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" validate:"required,min=5,max=45"`
}

//CreateToCategory To category
func CreateToCategory(categoryDto CreateCategoryDto) models.Category {
	return models.Category{Name: categoryDto.Name, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
}
