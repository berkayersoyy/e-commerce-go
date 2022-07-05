package user

import (
	"errors"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers/dto"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/twinj/uuid"
	"net/http"
	"time"
)

//userHandler User handler
type userHandler struct {
	userService services.UserService
}

func (u userHandler) Insert(c *gin.Context) {
	var user dto.CreateUserDto
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	userToAdd := models.User{
		UUID:      uuid.NewV4().String(),
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	err = u.userService.Insert(c, userToAdd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (u userHandler) FindByUUID(c *gin.Context) {
	id := c.Param("uuid")
	user, err := u.userService.FindByUUID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

func (u userHandler) FindByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := u.userService.FindByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errors.New(models.UserNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

func (u userHandler) Update(c *gin.Context) {
	var userDto dto.UpdateUserDto
	err := c.BindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user, err := u.userService.FindByUUID(c, userDto.UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errors.New("user not found")})
		return
	}
	user.Username = userDto.Username
	user.Password = userDto.Password
	err = u.userService.Update(c, user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusCreated)
}

func (u userHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	err := u.userService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

//ProvideUserHandler Provide user handler dynamodb
func ProvideUserHandler(u services.UserService) handlers.UserHandler {
	return userHandler{userService: u}
}
