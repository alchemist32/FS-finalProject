package repository

import "github.com/products-api/products/domain/models"

type IProductRepository interface {
	GetAll() ([]*models.Product, error)
	GetByCodeBar(codeBar string) (*models.Product, error)
	Create(product *models.Product) error
}
