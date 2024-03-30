package repository

import "github.com/FS-finalProyect/productss-api/products/domain/models"

type IProductRepository interface {
	GetAll() ([]*models.Product, error)
	GetByCodeBar(codeBar string) (*models.Product, error)
	Create(product *models.Product) error
}
