package handler

import (
	"github.com/products-api/core/infrastructure/adapter/client/mock"
	"github.com/products-api/products/application/usecases"
	"github.com/products-api/products/infraesturucture/dao"
)

func BuildProductsHandler() IProductHandler {
	dbClient := mock.NewMockClient()
	productHandler := NewProductsHandler(
		usecases.NewGetAllProducts(*dao.NewProductsDAO(dbClient)),
		usecases.NewAddProducts(*dao.NewProductsDAO(dbClient)),
	)

	return productHandler
}
