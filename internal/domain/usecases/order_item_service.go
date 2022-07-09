package usecases

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

type OrderItemService interface {
	GetAllOrderItems(c context.Context) []models.OrderItem
	GetOrderItemByID(c context.Context, id uint) models.OrderItem
	CreateOrderItem(c context.Context, orderItem models.OrderItem) models.OrderItem
	DeleteOrderItem(c context.Context, orderItem models.OrderItem)
	GetOrderItemsByOrderID(c context.Context, id uint) []models.OrderItem
}
