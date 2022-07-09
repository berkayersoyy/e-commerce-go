package models

type OrderItem struct {
	ID        uint `gorm:"primary_key"`
	ProductId uint `json:"ProductId" validate:"required"`
	OrderId   uint `json:"OrderId"`
	Amount    int  `json:"Amount" validate:"required"`
}
