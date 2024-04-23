package usecases

import (
	"errors"

	"github.com/products-api/products/domain/models"
	"github.com/products-api/products/domain/repository"
)

type IGetProductByBarcode interface {
	Execute(barcode string) (*models.Product, error)
}

type getProductByBarcode struct {
	productRepo repository.ProductsRepository
}

// NewGetProductByBarcode initialize the use case
func NewGetProductByBarcode(productRepo repository.ProductsRepository) *getProductByBarcode {
	return &getProductByBarcode{
		productRepo: productRepo,
	}
}

// Execute retrieves a product
// It will return a product struct or an error y something went wrong
func (pdao *getProductByBarcode) Execute(barcode string) (*models.Product, error) {
	product, err := pdao.productRepo.GetProductByBarCode(barcode)

	if err != nil && errors.Is(repository.NotFoundProduct, err) {
		return nil, repository.NotFoundProduct
	}

	if err != nil {
		return nil, err
	}
	return product, nil
}
