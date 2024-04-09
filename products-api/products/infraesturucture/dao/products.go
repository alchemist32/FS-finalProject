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
	item["id"] = product.ID
	item["name"] = product.Name
	item["description"] = product.Description
	item["price"] = product.Price
	item["barcode"] = product.BarCode
	item["stock"] = product.Stock
	_, err := pdao.Client.CreateItem(item)
	if err != nil {
		return ErrorAddProduct
	}
	return nil
}

func (pdao ProductsDAO) GetProductByBarCode(barcode string) (*models.Product, error) {
	var product models.Product
	result, err := pdao.Client.GetItemByBarcode(barcode)

	if err != nil && errors.Is(mock.NotFoundDBItem, err) {
		return nil, NotFoundProduct
	}
	errUnmarShal := json.Unmarshal(result, &product)

	if errUnmarShal != nil {
		return nil, UnmarshalError
	}

	return &product, nil
}

func (pdao ProductsDAO) GetProductStockById(productId int) int {
	return pdao.Client.GetStock(productId)
}
