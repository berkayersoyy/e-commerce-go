package handlers

import (
	"github.com/gin-gonic/gin"
)

//ProductHandler Product handler
type ProductHandler interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetProductsByCategoryID(c *gin.Context)
}
