package repositories

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
)

type ProductRepository interface {
	GetAllProducts(c context.Context) []models.Product
	GetProductByID(c context.Context, id uint) models.Product
	AddProduct(c context.Context, product models.Product) models.Product
	DeleteProduct(c context.Context, product models.Product)
}