package mock

import (
	"errors"
	"fmt"
)

var (
	SuccessDBMSG     = "connected successfully"
	NotFoundDBItem   = errors.New("error: db item was not found")
	AlreadyExistItem = errors.New("the Id already exist in db")
)

type IMockClient interface {
	GetItems() *[]map[string]any
	GetItem(itemId int) (*map[string]any, error)
	GetItemByBarcode(barcode string) (*map[string]any, error)
	CreateItem(item map[string]any) (int, error)
	InitMockClient() (string, error)
}

type MockClient struct {
	data []map[string]any
}

func NewMockClient() *MockClient {
	mc := MockClient{}
	mc.InitMockClient()
	return &mc
}

func (mc MockClient) InitMockClient() (string, error) {
	products := []map[string]any{
		{
			"id":          1,
			"name":        "Product 1",
			"description": "Description of product 1",
			"barcode":     "1234567890",
			"price":       10.99,
		},
		{
			"id":          2,
			"name":        "Product 2",
			"description": "Description of product 2",
			"barcode":     "0987654321",
			"price":       19.99,
		},
		{
			"id":          3,
			"name":        "Product 3",
			"description": "Description of product 3",
			"barcode":     "5678901234",
			"price":       5.99,
		},
		{
			"id":          4,
			"name":        "Product 4",
			"description": "Description of product 4",
			"barcode":     "4321098765",
			"price":       7.99,
		},
		{
			"id":          5,
			"name":        "Product 5",
			"description": "Description of product 5",
			"barcode":     "6789012345",
			"price":       14.99,
		},
	}
	mc.data = products
	fmt.Println(SuccessDBMSG)
	return SuccessDBMSG, nil
}

func (mc MockClient) GetItems() *[]map[string]any {
	return &mc.data
}

func (mc MockClient) GetItemByBarcode(barcode string) (*map[string]any, error) {
	var p *map[string]any
	for _, product := range mc.data {
		item := product["barcode"]
		if item == barcode {
			p = &product
		}
	}

	if p == nil {
		return nil, NotFoundDBItem
	}
	return p, nil
}
