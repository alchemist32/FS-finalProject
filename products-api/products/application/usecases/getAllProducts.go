package usecases

import (
	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/infraesturucture/dao"
)

type IGetAllProducts interface {
	Execute() (*[]models.Product, error)
}

type getAllProducts struct {
	productsDAO dao.ProductsDAO
}

func NewGetAllProducts(productsDAO dao.ProductsDAO) *getAllProducts {
	return &getAllProducts{
		productsDAO: productsDAO,
	}
}

func (p *getAllProducts) Execute() (*[]models.Product, error) {
	return p.productsDAO.GetAllProducts()
}
