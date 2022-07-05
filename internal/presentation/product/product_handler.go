package product

import (
	"errors"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers/dto"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//productHandler Product handler
type productHandler struct {
	ProductService services.ProductService
}

func (p *productHandler) GetAllProducts(c *gin.Context) {
	products := p.ProductService.GetAllProducts(c)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (p *productHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.ProductNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (p *productHandler) AddProduct(c *gin.Context) {
	var productDto dto.CreateProductDto
	err := c.BindJSON(&productDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	productToAdd := dto.CreateToProduct(productDto)
	createdProduct := p.ProductService.AddProduct(c, productToAdd)
	c.JSON(http.StatusCreated, gin.H{"product": createdProduct})
}

func (p *productHandler) UpdateProduct(c *gin.Context) {
	var productDto dto.UpdateProductDto
	err := c.BindJSON(&productDto)
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
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.ProductNotFound)})
		return
	}

	product.Name = productDto.Name
	product.Price = productDto.Price
	product.Description = productDto.Description
	p.ProductService.AddProduct(c, product)
	c.Status(http.StatusCreated)
}

func (p *productHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.ProductNotFound)})
		return
	}

	p.ProductService.DeleteProduct(c, product)
	c.Status(http.StatusCreated)
}

//ProvideProductHandler Provide product handler
func ProvideProductHandler(p services.ProductService) handlers.ProductHandler {
	return &productHandler{ProductService: p}
}
