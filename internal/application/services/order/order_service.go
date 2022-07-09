package order

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
)

//orderService Order service
type orderService struct {
	OrderRepository  repositories.OrderRepository
	ProductService   usecases.ProductService
	OrderItemService usecases.OrderItemService
}

func (o *orderService) GetAllOrders(c context.Context) []models.OrderDetail {
	var orderDetails []models.OrderDetail
	orders := o.OrderRepository.GetAllOrders(c)
	for _, order := range orders {
		orderItems := o.OrderItemService.GetOrderItemsByOrderID(c, order.ID)
		orderDetail := models.OrderDetail{
			OrderId:    order.ID,
			UserId:     order.UserId,
			TotalPrice: order.TotalPrice,
			OrderItems: orderItems,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return orderDetails
}

func (o *orderService) GetOrderByID(c context.Context, id uint) models.OrderDetail {
	order := o.OrderRepository.GetOrderByID(c, id)
	orderItems := o.OrderItemService.GetOrderItemsByOrderID(c, id)
	orderDetails := models.ToOrderDetail(order, orderItems)
	return orderDetails
}

func (o *orderService) GetOrdersByUserID(c context.Context, id uint) []models.OrderDetail {
	var orderDetails []models.OrderDetail
	orders := o.OrderRepository.GetOrdersByUserID(c, id)
	for _, order := range orders {
		orderItems := o.OrderItemService.GetOrderItemsByOrderID(c, order.ID)
		orderDetail := models.OrderDetail{
			OrderId:    order.ID,
			UserId:     order.UserId,
			TotalPrice: order.TotalPrice,
			OrderItems: orderItems,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return orderDetails
}

func (o *orderService) CreateOrder(c context.Context, order models.Order, orderItems []models.OrderItem) models.OrderDetail {
	order.TotalPrice = o.getTotalPrice(c, orderItems)
	addedOrder := o.OrderRepository.CreateOrder(c, order)
	o.addOrderItems(c, orderItems, addedOrder.ID)
	orderDetail := models.ToOrderDetail(addedOrder, orderItems)
	return orderDetail
}

func (o *orderService) DeleteOrder(c context.Context, order models.Order) {
	o.OrderRepository.DeleteOrder(c, order)
	o.deleteOrderItems(c, order.ID)
}

func (o *orderService) getTotalPrice(c context.Context, orderItems []models.OrderItem) float32 {
	var sumPrice float32
	for _, order := range orderItems {
		product := o.ProductService.GetProductByID(c, order.ProductId)
		price := float32(order.Amount) * product.Price
		sumPrice = sumPrice + price
	}
	return sumPrice
}

func (o *orderService) addOrderItems(c context.Context, orderItems []models.OrderItem, orderId uint) {
	for _, orderItem := range orderItems {
		orderItem.OrderId = orderId
		o.OrderItemService.CreateOrderItem(c, orderItem)
	}
}

func (o *orderService) deleteOrderItems(c context.Context, orderId uint) {
	orderItems := o.OrderItemService.GetOrderItemsByOrderID(c, orderId)
	for _, orderItem := range orderItems {
		o.OrderItemService.DeleteOrderItem(c, orderItem)
	}
}

// ProvideOrderService Provide order service
func ProvideOrderService(o repositories.OrderRepository, p usecases.ProductService, ot usecases.OrderItemService) usecases.OrderService {
	return &orderService{OrderRepository: o, ProductService: p, OrderItemService: ot}
}
