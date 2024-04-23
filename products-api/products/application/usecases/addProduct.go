package usecases

import (
	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/domain/repository"
)

type IAddProducts interface {
	Execute(product models.Product) error
}

type addProduct struct {
	productsDAO repository.ProductsRepository
}

// NewAddProducts initialize te use case
func NewAddProducts(productsDAO repository.ProductsRepository) *addProduct {
	return &addProduct{
		productsDAO: productsDAO,
	}
}

// Execute add a new product
func (p *addProduct) Execute(product models.Product) error {
	return p.productsDAO.AddProduct(product)
}
