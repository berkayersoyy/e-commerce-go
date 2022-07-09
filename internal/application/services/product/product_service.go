package product

import (
	"context"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/models"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/repositories"
	"github.com/berkayersoyy/e-commerce-go/internal/domain/usecases"
)

//productService Product service
type productService struct {
	ProductRepository repositories.ProductRepository
}

func (p *productService) GetAllProducts(c context.Context) []models.Product {
	return p.ProductRepository.GetAllProducts(c)
}

func (p *productService) GetProductByID(c context.Context, id uint) models.Product {
	return p.ProductRepository.GetProductByID(c, id)
}

func (p *productService) GetProductsByCategoryID(c context.Context, id uint) []models.Product {
	return p.ProductRepository.GetProductsByCategoryID(c, id)
}

func (p *productService) CreateProduct(c context.Context, product models.Product) models.Product {
	p.ProductRepository.CreateProduct(c, product)
	return product
}

func (p *productService) DeleteProduct(c context.Context, product models.Product) {
	p.ProductRepository.DeleteProduct(c, product)
}

// ProvideProductService Provide product service
func ProvideProductService(p repositories.ProductRepository) usecases.ProductService {
	return &productService{ProductRepository: p}
}
