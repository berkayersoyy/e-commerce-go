package order_item

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/jinzhu/gorm"
)

//orderItemRepository Order item repository
type orderItemRepository struct {
	DB *gorm.DB
}

func (o *orderItemRepository) GetAllOrderItems(c context.Context) []models.OrderItem {
	var orderItems []models.OrderItem
	o.DB.Find(&orderItems)
	return orderItems
}

func (o *orderItemRepository) GetOrderItemByID(c context.Context, id uint) models.OrderItem {
	var orderItem models.OrderItem
	o.DB.First(&orderItem, id)
	return orderItem
}

func (o *orderItemRepository) GetOrderItemsByOrderID(c context.Context, id uint) []models.OrderItem {
	var orderItems []models.OrderItem
	o.DB.Where("order_id = ?", id).Find(&orderItems)
	return orderItems
}

func (o *orderItemRepository) CreateOrderItem(c context.Context, orderItem models.OrderItem) models.OrderItem {
	o.DB.Save(&orderItem)
	return orderItem
}

func (o *orderItemRepository) DeleteOrderItem(c context.Context, orderItem models.OrderItem) {
	o.DB.Delete(&orderItem)
}

//ProvideOrderItemRepository Provide order item repository
func ProvideOrderItemRepository(db *gorm.DB) repositories.OrderItemRepository {
	return &orderItemRepository{DB: db}
}
