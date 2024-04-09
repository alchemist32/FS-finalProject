package usecases

import "github.com/products-api/products/domain/repository"

type IGetProductStockById interface {
	Execute(productId int, channel chan int)
}

type getProductStockById struct {
	productRepo repository.ProductsRepository
}

func NewGetProductStockById(productRepo repository.ProductsRepository) *getProductStockById {
	return &getProductStockById{
		productRepo: productRepo,
	}
}

func (pdao *getProductStockById) Execute(productId int, channel chan int) {
	channel <- pdao.productRepo.GetProductStockById(productId)
}
