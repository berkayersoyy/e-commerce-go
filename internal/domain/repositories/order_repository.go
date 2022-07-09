package repositories

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

type OrderRepository interface {
	GetAllOrders(c context.Context) []models.Order
	GetOrderByID(c context.Context, id uint) models.Order
	CreateOrder(c context.Context, order models.Order) models.Order
	DeleteOrder(c context.Context, order models.Order)
	GetOrdersByUserID(c context.Context, id uint) []models.Order
}
