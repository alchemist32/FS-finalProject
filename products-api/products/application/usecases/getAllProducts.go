package usecases

import (
	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/domain/repository"
)

type IGetAllProducts interface {
	Execute() (*[]models.Product, error)
}

type getAllProducts struct {
	productsDAO repository.ProductsRepository
}

// NewGetAllProducts initialize the use case
func NewGetAllProducts(productsDAO repository.ProductsRepository) *getAllProducts {
	return &getAllProducts{
		productsDAO: productsDAO,
	}
}

// Execute retrieves the list of products
func (p *getAllProducts) Execute() (*[]models.Product, error) {
	return p.productsDAO.GetAllProducts()
}
