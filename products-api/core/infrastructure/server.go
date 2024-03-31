package infrastructure

import (
	"github.com/products-api/core/infrastructure/adapter/client/mock"
)

func Setup() {
	mock.NewMockClient()
}
