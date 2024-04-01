package dao

import (
	"encoding/json"
	"errors"

	"github.com/products-api/core/infrastructure/adapter/client/mock"
	"github.com/products-api/products/domain/models"
)

type ProductsDAO struct {
	Client mock.IMockClient
}

var (
	NotFoundProducts = errors.New("no products on db")
	NotFoundProduct  = errors.New("product not found")
	UnmarshalError   = errors.New("json could not be unmarshal")
)

func NewProductsDAO(client mock.IMockClient) *ProductsDAO {
	return &ProductsDAO{
		Client: client,
	}
}

func (pdao ProductsDAO) GetAllProducts() (*[]models.Product, error) {
	var result []models.Product

	err := json.Unmarshal(pdao.Client.GetItems(), &result)

	if err != nil {
		return nil, UnmarshalError
	}
	return &result, nil
}
