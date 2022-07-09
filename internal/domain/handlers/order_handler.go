package handlers

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	GetAllOrders(c *gin.Context)
	GetOrderByID(c *gin.Context)
	CreateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
	GetOrdersByUserID(c *gin.Context)
	UpdateOrder(c *gin.Context)
}
