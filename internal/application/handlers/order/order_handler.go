package order

import (
	"errors"
	"github.com/berkayersoyy/e-commerce-go/internal/application/utils"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers/dto"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//orderHandler Order handler
type orderHandler struct {
	OrderService usecases.OrderService
}

// @BasePath /api/v1

// GetAllOrders
// @Summary Fetch all order
// @Schemes
// @Description Fetch all order
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} models.OrderDetail
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/orders/ [get]
func (p *orderHandler) GetAllOrders(c *gin.Context) {
	orders := p.OrderService.GetAllOrders(c)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// @BasePath /api/v1

// GetOrderByID
// @Summary Fetch order by id
// @Schemes
// @Description Fetch order by id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.OrderDetail
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/orders/{id} [get]
func (p *orderHandler) GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order := p.OrderService.GetOrderByID(c, uint(id))
	if order.OrderItems == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.OrderNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order": order})
}

// @BasePath /api/v1

// GetOrdersByUserID
// @Summary Fetch order by user id
// @Schemes
// @Description Fetch order by user id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.OrderDetail
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/orders/getbyuserid/{id} [get]
func (p *orderHandler) GetOrdersByUserID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	orders := p.OrderService.GetOrdersByUserID(c, uint(id))
	if orders == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.OrderNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})

}

// @BasePath /api/v1

// CreateOrder
// @Summary Add Order
// @Schemes
// @Description Add Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param product body dto.CreateOrderDto true "Order"
// @Success 200 {object} models.Order
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/orders/ [post]
func (p *orderHandler) CreateOrder(c *gin.Context) {
	var orderDto dto.CreateOrderDto
	err := c.BindJSON(&orderDto)
	if err != nil { //TODO dev ya da proda gore error
		c.JSON(utils.HttpBadRequestError(err, err.Error()))
		return
	}
	validate := validator.New()
	err = validate.Struct(orderDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	orderToAdd := dto.CreateToOrder(orderDto)
	p.OrderService.CreateOrder(c, orderToAdd, orderDto.OrderItems)
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// UpdateOrder
// @Summary Update Order
// @Schemes
// @Description Update Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param product body dto.UpdateOrderDto true "Order Dto"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/orders/ [put]
func (p *orderHandler) UpdateOrder(c *gin.Context) {
	var orderDto dto.UpdateOrderDto
	err := c.BindJSON(&orderDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	order := p.OrderService.GetOrderByID(c, uint(id))
	if order.OrderItems == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.OrderNotFound)})
		return
	}

	orderToUpdated := dto.UpdateToOrder(orderDto)
	p.OrderService.CreateOrder(c, orderToUpdated, orderDto.OrderItems)
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// DeleteOrder
// @Summary Delete Order
// @Schemes
// @Description Delete Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/orders/{id} [delete]
func (p *orderHandler) DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order := p.OrderService.GetOrderByID(c, uint(id))
	if order.OrderItems == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.OrderNotFound)})
		return
	}
	orderToDeleted := models.ToOrder(order)
	p.OrderService.DeleteOrder(c, orderToDeleted)
	c.Status(http.StatusCreated)
}

//ProvideOrderHandler Provide order handler
func ProvideOrderHandler(o usecases.OrderService) handlers.OrderHandler {
	return &orderHandler{OrderService: o}
}
