package handlers

import (
	"github.com/gin-gonic/gin"
)

//ProductHandler Product handler
type ProductHandler interface {
	GetAllProducts(c *gin.Context)
	GetProductByID(c *gin.Context)
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetProductByCategoryID(c *gin.Context)
}
