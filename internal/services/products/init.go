package products

import (
	"WeMarketOnGolang/internal/models"
	"context"
)

// ProductService описывает методы, которые должны быть реализованы для работы с продуктами
type ProductServiceInterface interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProduct(ctx context.Context, id int32) (*models.Product, error)
	ListProducts(ctx context.Context) ([]*models.Product, error)
	UpdateProduct(ctx context.Context, id int32, updatedData *models.Product) error
	DeleteProduct(ctx context.Context, id int32) error
}
