package category

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/jinzhu/gorm"
)

//categoryRepository Product repository
type categoryRepository struct {
	DB *gorm.DB
}

func (p *categoryRepository) GetAllCategories(c context.Context) []models.Category {
	var categories []models.Category
	p.DB.Find(&categories)
	return categories
}

func (p *categoryRepository) GetCategoryByID(c context.Context, id uint) models.Category {
	var category models.Category
	p.DB.First(&category, id)
	return category
}

func (p *categoryRepository) CreateCategory(c context.Context, category models.Category) models.Category {
	p.DB.Save(&category)
	return category
}

func (p *categoryRepository) DeleteCategory(c context.Context, category models.Category) {
	p.DB.Delete(&category)
}

//ProvideCategoryRepository Provide category repository
func ProvideCategoryRepository(db *gorm.DB) repositories.CategoryRepository {
	return &categoryRepository{DB: db}
}
