package category

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
)

//categoryService Category service
type categoryService struct {
	CategoryRepository repositories.CategoryRepository
}

func (p *categoryService) GetAllCategories(c context.Context) []models.Category {
	return p.CategoryRepository.GetAllCategories(c)
}

func (p *categoryService) GetCategoryByID(c context.Context, id uint) models.Category {
	return p.CategoryRepository.GetCategoryByID(c, id)
}

func (p *categoryService) CreateCategory(c context.Context, category models.Category) models.Category {
	p.CategoryRepository.CreateCategory(c, category)
	return category
}

func (p *categoryService) DeleteCategory(c context.Context, category models.Category) {
	p.CategoryRepository.DeleteCategory(c, category)
}

// ProvideCategoryService Provide category service
func ProvideCategoryService(c repositories.CategoryRepository) usecases.CategoryService {
	return &categoryService{CategoryRepository: c}
}
