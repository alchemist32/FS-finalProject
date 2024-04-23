package usecases

import "github.com/products-api/products/domain/repository"

type IGetProductStockById interface {
	Execute(productId int, channel chan int)
}

type getProductStockById struct {
	productRepo repository.ProductsRepository
}

// NewGetProductStockById initialize the use case
func NewGetProductStockById(productRepo repository.ProductsRepository) *getProductStockById {
	return &getProductStockById{
		productRepo: productRepo,
	}
}

// Execute retrieve the stock of a given product based on the id
func (pdao *getProductStockById) Execute(productId int, channel chan int) {
	channel <- pdao.productRepo.GetProductStockById(productId)
}
