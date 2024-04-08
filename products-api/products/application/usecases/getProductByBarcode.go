package usecases

import (
	"errors"

	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/infraesturucture/dao"
)

type IGetProductByBarcode interface {
	Execute(barcode string) (*models.Product, error)
}

type getProductByBarcode struct {
	productDAO dao.ProductsDAO
}

func NewGetProductByBarcode(productDAO dao.ProductsDAO) *getProductByBarcode {
	return &getProductByBarcode{
		productDAO: productDAO,
	}
}

func (pdao *getProductByBarcode) Execute(barcode string) (*models.Product, error) {
	product, err := pdao.productDAO.GetProductByBarCode(barcode)

	if err != nil && errors.Is(dao.NotFoundProduct, err) {
		return nil, dao.NotFoundProduct
	}

	if err != nil {
		return nil, err
	}
	return product, nil
}
