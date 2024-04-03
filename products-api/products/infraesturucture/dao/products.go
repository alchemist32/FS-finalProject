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
	ErrorAddProduct  = errors.New("Product cannot be added")
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

func (pdao ProductsDAO) AddProduct(product models.Product) error {
	var item map[string]any

	item = make(map[string]any)
	item["ID"] = product.ID
	item["Name"] = product.Name
	item["Description"] = product.Description
	item["Price"] = product.Price
	item["Barcode"] = product.BarCode
	item["Stock"] = product.Stock
	_, err := pdao.Client.CreateItem(item)
	if err != nil {
		return ErrorAddProduct
	}
	return nil
}
