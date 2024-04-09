package handler

import (
	"github.com/products-api/core/infrastructure/adapter/client/mock"
	"github.com/products-api/products/application/usecases"
	"github.com/products-api/products/domain/repository"
)

func BuildProductsHandler() IProductHandler {
	dbClient := mock.NewMockClient()
	productHandler := NewProductsHandler(
		usecases.NewGetAllProducts(*repository.NewProductsRepo(dbClient)),
		usecases.NewAddProducts(*repository.NewProductsRepo(dbClient)),
		usecases.NewGetProductByBarcode(*repository.NewProductsRepo(dbClient)),
		usecases.NewGetProductStockById(*repository.NewProductsRepo(dbClient)),
	)

	return productHandler
}
