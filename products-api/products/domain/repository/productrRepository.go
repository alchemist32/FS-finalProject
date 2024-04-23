package repository

import (
	"encoding/json"
	"errors"

	"github.com/products-api/core/infrastructure/adapter/client/mock"
	"github.com/products-api/products/domain/models"
)

type IProductRepository interface {
	GetAllProducts() (*[]models.Product, error)
	GetProductByBarCode(barcode string) (*models.Product, error)
	AddProduct(product models.Product) error
	GetProductStockById(productId int) int
}

type ProductsRepository struct {
	Client mock.IMockClient
}

var (
	NotFoundProducts = errors.New("no products on db")
	NotFoundProduct  = errors.New("product not found")
	UnmarshalError   = errors.New("json could not be unmarshal")
	ErrorAddProduct  = errors.New("Product cannot be added")
)

func NewProductsRepo(client mock.IMockClient) *ProductsRepository {
	return &ProductsRepository{
		Client: client,
	}
}

// GetAllProducts retrieves the list of products from the DB
func (pdao ProductsRepository) GetAllProducts() (*[]models.Product, error) {
	var result []models.Product

	err := json.Unmarshal(pdao.Client.GetItems(), &result)

	if err != nil {
		return nil, UnmarshalError
	}

	return &result, nil
}

// AddProduct using the provided argument add a new product to the list
func (pdao ProductsRepository) AddProduct(product models.Product) error {
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

// GetProductByBarCode retrieves the corresponding product using the barcode
func (pdao ProductsRepository) GetProductByBarCode(barcode string) (*models.Product, error) {
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

// GetProductStockById retrieves the stock for given product
func (pdao ProductsRepository) GetProductStockById(productId int) int {
	return pdao.Client.GetStock(productId)
}
