package handlers

import (
	"github.com/gin-gonic/gin"
)

//CategoryHandler Category service
type CategoryHandler interface {
	GetAllCategories(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	CreateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
}
