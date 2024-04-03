package usecases

import (
	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/infraesturucture/dao"
)

type IAddProducts interface {
	Execute(product models.Product) error
}

type addProduct struct {
	productsDAO dao.ProductsDAO
}

func NewAddProducts(productsDAO dao.ProductsDAO) *addProduct {
	return &addProduct{
		productsDAO: productsDAO,
	}
}

func (p *addProduct) Execute(product models.Product) error {
	return p.productsDAO.AddProduct(product)
}
