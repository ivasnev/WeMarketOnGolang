package products

import (
	"WeMarketOnGolang/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

// NewProductService создаёт новый экземпляр ProductService
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

// CreateProduct создаёт новый товар
func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	if err := s.DB.WithContext(ctx).Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProduct получает товар по ID
func (s *ProductService) GetProduct(ctx context.Context, id int32) (*models.Product, error) {
	var product models.Product
	if err := s.DB.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// ListProducts возвращает список всех товаров
func (s *ProductService) ListProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	if err := s.DB.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct обновляет информацию о товаре по ID
func (s *ProductService) UpdateProduct(ctx context.Context, id int32, updatedData *models.Product) error {
	var product models.Product
	// Ищем товар по ID
	if err := s.DB.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	// Обновляем нужные поля
	product.Name = updatedData.Name
	product.Description = updatedData.Description
	product.Price = updatedData.Price
	// Применяем изменения
	if err := s.DB.WithContext(ctx).Save(&product).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct удаляет товар по ID
func (s *ProductService) DeleteProduct(ctx context.Context, id int32) error {
	if err := s.DB.WithContext(ctx).Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
