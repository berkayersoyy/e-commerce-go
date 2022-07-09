package dto

import (
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/jinzhu/gorm"
	"time"
)

type CreateOrderDto struct {
	UserId     string             `json:"UserId" validate:"required"`
	TotalPrice float32            `json:"TotalPrice"`
	OrderItems []models.OrderItem `json:"OrderItems" validate:"required"`
}

type UpdateOrderDto struct {
	UserId     string             `json:"UserId" validate:"required"`
	TotalPrice float32            `json:"TotalPrice"`
	OrderItems []models.OrderItem `json:"OrderItems" validate:"required"`
}

func CreateToOrder(dto CreateOrderDto) models.Order {
	return models.Order{UserId: dto.UserId, TotalPrice: dto.TotalPrice, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil}}
}
func UpdateToOrder(dto UpdateOrderDto) models.Order {
	return models.Order{UserId: dto.UserId, TotalPrice: dto.TotalPrice, Model: gorm.Model{UpdatedAt: time.Now(), DeletedAt: nil}}
}
