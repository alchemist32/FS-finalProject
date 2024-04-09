package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	SuccessDBMSG     = "connected successfully"
	NotFoundDBItem   = errors.New("error: db item was not found")
	AlreadyExistItem = errors.New("the Id already exist in db")
	CannotAddItem    = errors.New("the item cannot be added")
)

type IMockClient interface {
	GetItems() []byte
	// GetItem(itemId int) ([]byte, error)
	GetItemByBarcode(barcode string) ([]byte, error)
	CreateItem(item map[string]any) (int, error)
	GetStock(itemId int) int
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

func (mc *MockClient) InitMockClient() (string, error) {
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

func (mc *MockClient) GetItems() []byte {
	jsonString, _ := json.Marshal(mc.data)
	return jsonString
}

func (mc MockClient) GetItemByBarcode(barcode string) ([]byte, error) {
	var p *map[string]any
	for _, product := range mc.data {
		item := product["barcode"]
		if item == barcode {
			p = &product
			break
		}
	}

	if p == nil {
		return nil, NotFoundDBItem
	}

	jString, _ := json.Marshal(p)
	return jString, nil
}

func (mc *MockClient) CreateItem(item map[string]any) (int, error) {
	itemsSlice := mc.data
	currentLen := len(itemsSlice)
	item["id"] = currentLen + 1
	newState := append(itemsSlice, item)
	dataLen := len(newState)
	if currentLen == dataLen {
		return 0, CannotAddItem
	}
	mc.data = newState
	return dataLen - 1, nil
}

func (mc *MockClient) GetStock(itemId int) int {
	time.Sleep(300)
	maxNumber := itemId * 100
	return rand.Intn(maxNumber)
}
