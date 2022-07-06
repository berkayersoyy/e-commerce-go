package product

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

//productRepository Product repository
type productRepository struct {
	DB *gorm.DB
}

func (p *productRepository) GetAllProducts(c context.Context) []models.Product {
	var products []models.Product
	p.DB.Find(&products)
	return products
}

func (p *productRepository) GetProductByID(c context.Context, id uint) models.Product {
	var product models.Product
	p.DB.First(&product, id)
	return product
}

func (p *productRepository) AddProduct(c context.Context, product models.Product) models.Product {
	p.DB.Save(&product)
	return product
}

func (p *productRepository) DeleteProduct(c context.Context, product models.Product) {
	p.DB.Delete(&product)
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

	db.AutoMigrate(&models.Product{})

	return db
}

//ProvideProductRepository Provide product repository
func ProvideProductRepository() repositories.ProductRepository {
	return &productRepository{DB: initDb()}
}
