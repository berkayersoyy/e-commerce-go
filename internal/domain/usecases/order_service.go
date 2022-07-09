package usecases

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

type OrderService interface {
	GetAllOrders(c context.Context) []models.OrderDetail
	GetOrderByID(c context.Context, id uint) models.OrderDetail
	CreateOrder(c context.Context, order models.Order, orderItems []models.OrderItem) models.OrderDetail
	DeleteOrder(c context.Context, order models.Order)
	GetOrdersByUserID(c context.Context, id uint) []models.OrderDetail
}
