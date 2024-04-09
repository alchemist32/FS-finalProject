package usecases

import "github.com/products-api/products/infraesturucture/dao"

type IGetProductStockById interface {
	Execute(productId int, channel chan int)
}

type getProductStockById struct {
	productDAO dao.ProductsDAO
}

func NewGetProductStockById(productDAO dao.ProductsDAO) *getProductStockById {
	return &getProductStockById{
		productDAO: productDAO,
	}
}

func (pdao *getProductStockById) Execute(productId int, channel chan int) {
	channel <- pdao.productDAO.GetProductStockById(productId)
}
