package user

import (
	"errors"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers/dto"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//userHandler User handler
type userHandler struct {
	userService services.UserService
}

// @BasePath /api/v1

// Insert
// @Summary Add user
// @Schemes
// @Description Add user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User"
// @Success 200 {object} models.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/ [post]
func (u userHandler) Insert(c *gin.Context) {
	var userDto dto.CreateUserDto
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
	userToAdd := dto.ToUser(userDto)
	err = u.userService.Insert(c, userToAdd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// FindByUUID
// @Summary Find user
// @Schemes
// @Description Find user by uuid
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} models.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/getbyuuid/{uuid} [get]
func (u userHandler) FindByUUID(c *gin.Context) {
	id := c.Param("uuid")
	user, err := u.userService.FindByUUID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": models.UserNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// @BasePath /api/v1

// FindByUsername
// @Summary Find user
// @Schemes
// @Description Find user by username
// @Tags Users
// @Accept json
// @Produce json
// @Param username path string true "User Username"
// @Success 200 {object} models.User
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/getbyusername/{username} [get]
func (u userHandler) FindByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := u.userService.FindByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if user == (models.User{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.UserNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// @BasePath /api/v1

// Update
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User Dto"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/ [put]
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
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.UserNotFound)})
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

// @BasePath /api/v1

// Delete
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /v1/users/{id} [delete]
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
