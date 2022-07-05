package dto

import (
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/jinzhu/gorm"
	"time"
)

//UpdateProductDto Update product dto
type UpdateProductDto struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Name        string  `json:"name" validate:"required,min=5,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

//CreateProductDto Create product dto
type CreateProductDto struct {
	Name        string  `json:"name" validate:"required,min=5,max=45"`
	Price       float32 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required"`
}

//CreateToProduct To product
func CreateToProduct(productDto CreateProductDto) models.Product {
	return models.Product{Name: productDto.Name, Price: productDto.Price, Description: productDto.Description, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
}

//UpdateToProduct To_Product
func UpdateToProduct(productDto UpdateProductDto) models.Product {
	return models.Product{Name: productDto.Name, Price: productDto.Price, Description: productDto.Description, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
}
