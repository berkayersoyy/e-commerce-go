package models

import "github.com/jinzhu/gorm"

const (
	OrderNotFound = "product not found"
)

type Order struct {
	gorm.Model
	UserId     string  `json:"UserId" validate:"required"`
	TotalPrice float32 `json:"TotalPrice" validate:"required"`
}

type OrderDetail struct {
	OrderId    uint        `json:"OrderId" gorm:"primary_key"`
	UserId     string      `json:"UserId" validate:"required"`
	TotalPrice float32     `json:"TotalPrice" validate:"required"`
	OrderItems []OrderItem `json:"OrderItems" validate:"required"`
}

func ToOrderDetail(order Order, orderItems []OrderItem) OrderDetail {
	orderDetail := OrderDetail{
		OrderId:    order.ID,
		UserId:     order.UserId,
		TotalPrice: order.TotalPrice,
		OrderItems: orderItems,
	}
	return orderDetail
}

func ToOrder(orderDetail OrderDetail) Order {
	return Order{
		UserId:     orderDetail.UserId,
		TotalPrice: orderDetail.TotalPrice,
		Model: gorm.Model{
			ID: orderDetail.OrderId,
		},
	}
}
