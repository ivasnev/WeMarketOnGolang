package products

import (
	"WeMarketOnGolang/internal/models"
)

// ProductService описывает методы, которые должны быть реализованы для работы с продуктами
type ProductServiceInterface interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id int32) (*models.Product, error)
	GetAllProducts() ([]*models.Product, error)
	UpdateProduct(id int32, updatedData *models.Product) error
	DeleteProduct(id int32) error
}
