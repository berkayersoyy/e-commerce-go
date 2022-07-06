package category

import (
	"context"
	"fmt"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/jinzhu/gorm"
	"github.com/sethvargo/go-retry"
	"log"
	"os"
	"time"
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

func (p *categoryRepository) AddCategory(c context.Context, category models.Category) models.Category {
	p.DB.Save(&category)
	return category
}

func (p *categoryRepository) DeleteCategory(c context.Context, category models.Category) {
	p.DB.Delete(&category)
}

//initDb Init db
func initDb() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	ctx := context.Background()
	var db *gorm.DB
	var err error
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println(err)
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	db.AutoMigrate(&models.Category{})

	return db
}

//ProvideCategoryRepository Provide category repository
func ProvideCategoryRepository() repositories.CategoryRepository {
	return &categoryRepository{DB: initDb()}
}
