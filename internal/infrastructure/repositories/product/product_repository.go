package product

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/jinzhu/gorm"
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

func (p *productRepository) GetProductsByCategoryID(c context.Context, id uint) []models.Product {
	var products []models.Product
	p.DB.Where("category_id = ?", id).Find(&products)
	return products
}

func (p *productRepository) CreateProduct(c context.Context, product models.Product) models.Product {
	p.DB.Save(&product)
	return product
}

func (p *productRepository) DeleteProduct(c context.Context, product models.Product) {
	p.DB.Delete(&product)
}

//ProvideProductRepository Provide product repository
func ProvideProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &productRepository{DB: db}
}
