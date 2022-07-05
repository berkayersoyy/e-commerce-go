package handlers

import "github.com/gin-gonic/gin"

//UserHandler User handler
type UserHandler interface {
	Update(c *gin.Context)
	FindByUUID(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
	FindByUsername(c *gin.Context)
}
