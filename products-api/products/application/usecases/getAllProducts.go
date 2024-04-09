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

func NewGetAllProducts(productsDAO repository.ProductsRepository) *getAllProducts {
	return &getAllProducts{
		productsDAO: productsDAO,
	}
}

func (p *getAllProducts) Execute() (*[]models.Product, error) {
	return p.productsDAO.GetAllProducts()
}
