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

// @BasePath /api/v1

// GetAllProducts
// @Summary Fetch all product
// @Schemes
// @Description Fetch all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [get]
func (p *productHandler) GetAllProducts(c *gin.Context) {
	products := p.ProductService.GetAllProducts(c)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// @BasePath /api/v1

// GetProductByID
// @Summary Fetch product by id
// @Schemes
// @Description Fetch product by id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/{id} [get]
func (p *productHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.GetProductByID(c, uint(id))
	if product == (models.Product{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.ProductNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

// @BasePath /api/v1

// AddProduct
// @Summary Add Product
// @Schemes
// @Description Add Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 200 {object} dto.CreateProductDto
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [post]
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

// @BasePath /api/v1

// UpdateProduct
// @Summary Update Product
// @Schemes
// @Description Update Product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.UpdateProductDto true "Product Dto"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/ [put]
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

// @BasePath /api/v1

// DeleteProduct
// @Summary Delete Product
// @Schemes
// @Description Delete Product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/products/{id} [delete]
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
