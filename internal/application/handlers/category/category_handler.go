package category

import (
	"errors"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/handlers/dto"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//categoryHandler Category handler
type categoryHandler struct {
	CategoryService usecases.CategoryService
}

// @BasePath /api/v1

// GetAllCategories
// @Summary Fetch all categories
// @Schemes
// @Description Fetch all categories
// @Tags Categories
// @Accept json
// @Produce json
// @Success 200 {object} models.Category
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/categories/ [get]
func (p *categoryHandler) GetAllCategories(c *gin.Context) {
	categories := p.CategoryService.GetAllCategories(c)
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// @BasePath /api/v1

// GetCategoryByID
// @Summary Fetch category by id
// @Schemes
// @Description Fetch category by id
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/categories/{id} [get]
func (p *categoryHandler) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category := p.CategoryService.GetCategoryByID(c, uint(id))
	if category == (models.Category{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.CategoryNotFound)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category})
}

// @BasePath /api/v1

// CreateCategory
// @Summary Add Category
// @Schemes
// @Description Add Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryDto true "Category"
// @Success 200 {object} models.Category
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/categories/ [post]
func (p *categoryHandler) CreateCategory(c *gin.Context) {
	var categoryDto dto.CreateCategoryDto
	err := c.BindJSON(&categoryDto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(categoryDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	categoryToAdd := dto.CreateToCategory(categoryDto)
	createdProduct := p.CategoryService.CreateCategory(c, categoryToAdd)
	c.JSON(http.StatusCreated, gin.H{"product": createdProduct})
}

// @BasePath /api/v1

// UpdateCategory
// @Summary Update Category
// @Schemes
// @Description Update Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body dto.UpdateCategoryDto true "Category Dto"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/categories/ [put]
func (p *categoryHandler) UpdateCategory(c *gin.Context) {
	var categoryDto dto.UpdateCategoryDto
	err := c.BindJSON(&categoryDto)
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
	category := p.CategoryService.GetCategoryByID(c, uint(id))
	if category == (models.Category{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.CategoryNotFound)})
		return
	}

	category.Name = categoryDto.Name
	p.CategoryService.CreateCategory(c, category)
	c.Status(http.StatusCreated)
}

// @BasePath /api/v1

// DeleteCategory
// @Summary Delete Category
// @Schemes
// @Description Delete Category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Security bearerAuth
// @Router /v1/categories/{id} [delete]
func (p *categoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category := p.CategoryService.GetCategoryByID(c, uint(id))
	if category == (models.Category{}) {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.New(models.ProductNotFound)})
		return
	}

	p.CategoryService.DeleteCategory(c, category)
	c.Status(http.StatusCreated)
}

//ProvideCategoryHandler Provide category handler
func ProvideCategoryHandler(c usecases.CategoryService) handlers.CategoryHandler {
	return &categoryHandler{CategoryService: c}
}
