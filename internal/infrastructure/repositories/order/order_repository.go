package order

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/jinzhu/gorm"
)

//orderRepository Order repository
type orderRepository struct {
	DB *gorm.DB
}

func (o *orderRepository) GetAllOrders(c context.Context) []models.Order {
	var orders []models.Order
	o.DB.Find(&orders)
	return orders
}

func (o *orderRepository) GetOrderByID(c context.Context, id uint) models.Order {
	var order models.Order
	o.DB.First(&order, id)
	return order
}

func (o *orderRepository) GetOrdersByUserID(c context.Context, id uint) []models.Order {
	var orders []models.Order
	o.DB.Where("user_id = ?", id).Find(&orders)
	return orders
}

func (o *orderRepository) CreateOrder(c context.Context, order models.Order) models.Order {
	o.DB.Save(&order)
	return order
}

func (o *orderRepository) DeleteOrder(c context.Context, order models.Order) {
	o.DB.Delete(&order)
}

//ProvideOrderRepository Provide order repository
func ProvideOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderRepository{DB: db}
}
