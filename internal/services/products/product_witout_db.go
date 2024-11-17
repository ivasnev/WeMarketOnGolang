package products

import (
	"WeMarketOnGolang/internal/models"
	"context"
	"errors"
	"sync"
	"time"
)

// InMemoryProductService — реализация ProductService для хранения данных в памяти
type InMemoryProductService struct {
	products map[int32]*models.Product
	nextID   int32
	mu       sync.Mutex
}

// SeedProducts заполняет InMemoryProductService начальными данными
func (s *InMemoryProductService) SeedProducts() {
	s.mu.Lock()
	defer s.mu.Unlock()

	var description1 = "Описание продукта 1"
	var imageURL1 = "http://example.com/image1.jpg"
	var sku1 = "SKU1"
	var weight1 float64 = 1.5
	var dimensions1 = "10x10x10"
	var manufacturerID1 int32 = 1

	var description2 = "Описание продукта 2"
	var imageURL2 = "http://example.com/image2.jpg"
	var sku2 = "SKU2"
	var weight2 float64 = 2.0
	var dimensions2 = "15x15x15"
	var manufacturerID2 int32 = 2

	products := []models.Product{
		{
			Name:               "Продукт 1",
			Description:        &description1,
			Price:              10.0,
			CategoryID:         1,
			Stock:              100,
			ImageURL:           &imageURL1,
			Options:            nil,
			Sku:                &sku1,
			Weight:             &weight1,
			Dimensions:         &dimensions1,
			AvailabilityStatus: 1,
			ManufacturerID:     &manufacturerID1,
		},
		{
			Name:               "Продукт 2",
			Description:        &description2,
			Price:              20.0,
			CategoryID:         2,
			Stock:              200,
			ImageURL:           &imageURL2,
			Options:            nil,
			Sku:                &sku2,
			Weight:             &weight2,
			Dimensions:         &dimensions2,
			AvailabilityStatus: 2,
			ManufacturerID:     &manufacturerID2,
		},
		// Добавьте сюда другие продукты, если нужно
	}
	now := time.Now()
	for _, product := range products {
		product.ID = s.nextID
		product.AddedDate = &now
		s.products[s.nextID] = &product
		s.nextID++
	}
}

// NewInMemoryProductService создает новый InMemoryProductService
func NewInMemoryProductService() *InMemoryProductService {
	return &InMemoryProductService{
		products: make(map[int32]*models.Product),
		nextID:   1,
	}
}

// CreateProduct добавляет новый продукт в память
func (s *InMemoryProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Присваиваем ID и дату добавления новому продукту
	product.ID = s.nextID
	now := time.Now()
	product.AddedDate = &now
	s.products[s.nextID] = product
	s.nextID++
	return nil
}

// GetProduct получает продукт по ID
func (s *InMemoryProductService) GetProduct(ctx context.Context, id int32) (*models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, exists := s.products[id]
	if !exists {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// ListProducts возвращает список всех продуктов
func (s *InMemoryProductService) ListProducts(ctx context.Context) ([]*models.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Копируем продукты из карты в срез
	products := make([]*models.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}
	return products, nil
}

// UpdateProduct обновляет продукт по ID
func (s *InMemoryProductService) UpdateProduct(ctx context.Context, id int32, updatedData *models.Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, exists := s.products[id]
	if !exists {
		return errors.New("product not found")
	}

	// Обновляем поля продукта
	product.Name = updatedData.Name
	product.Description = updatedData.Description
	product.Price = updatedData.Price
	product.CategoryID = updatedData.CategoryID
	product.Stock = updatedData.Stock
	product.ImageURL = updatedData.ImageURL
	product.Options = updatedData.Options
	product.Sku = updatedData.Sku
	product.Weight = updatedData.Weight
	product.Dimensions = updatedData.Dimensions
	product.AvailabilityStatus = updatedData.AvailabilityStatus
	product.ManufacturerID = updatedData.ManufacturerID

	return nil
}

// DeleteProduct удаляет продукт по ID
func (s *InMemoryProductService) DeleteProduct(ctx context.Context, id int32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.products[id]; !exists {
		return errors.New("product not found")
	}
	delete(s.products, id)
	return nil
}
