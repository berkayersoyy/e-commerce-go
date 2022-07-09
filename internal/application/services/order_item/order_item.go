package order_item

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
)

//orderItemService Order item repository
type orderItemService struct {
	OrderItemRepository repositories.OrderItemRepository
}

func (o *orderItemService) GetAllOrderItems(c context.Context) []models.OrderItem {
	return o.OrderItemRepository.GetAllOrderItems(c)
}

func (o *orderItemService) GetOrderItemByID(c context.Context, id uint) models.OrderItem {
	return o.OrderItemRepository.GetOrderItemByID(c, id)
}

func (o *orderItemService) GetOrderItemsByOrderID(c context.Context, id uint) []models.OrderItem {
	return o.OrderItemRepository.GetOrderItemsByOrderID(c, id)
}

func (o *orderItemService) CreateOrderItem(c context.Context, orderItem models.OrderItem) models.OrderItem {
	return o.OrderItemRepository.CreateOrderItem(c, orderItem)
}

func (o *orderItemService) DeleteOrderItem(c context.Context, orderItem models.OrderItem) {
	o.OrderItemRepository.DeleteOrderItem(c, orderItem)
}

//ProvideOrderItemService Provide order item repository
func ProvideOrderItemService(o repositories.OrderItemRepository) usecases.OrderItemService {
	return &orderItemService{OrderItemRepository: o}
}
