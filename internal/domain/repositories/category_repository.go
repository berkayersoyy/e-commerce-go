package repositories

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

//CategoryRepository Category repository
type CategoryRepository interface {
	GetAllCategories(c context.Context) []models.Category
	GetCategoryByID(c context.Context, id uint) models.Category
	CreateCategory(c context.Context, product models.Category) models.Category
	DeleteCategory(c context.Context, product models.Category)
}
