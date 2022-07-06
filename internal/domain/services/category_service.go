package services

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

//CategoryService Category service
type CategoryService interface {
	GetAllCategories(c context.Context) []models.Category
	GetCategoryByID(c context.Context, id uint) models.Category
	AddCategory(c context.Context, category models.Category) models.Category
	DeleteCategory(c context.Context, category models.Category)
}
