package repositories

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

//ProductRepository Product repository
type ProductRepository interface {
	GetAllProducts(c context.Context) []models.Product
	GetProductByID(c context.Context, id uint) models.Product
	CreateProduct(c context.Context, product models.Product) models.Product
	DeleteProduct(c context.Context, product models.Product)
	GetProductsByCategoryID(c context.Context, id uint) []models.Product
}
