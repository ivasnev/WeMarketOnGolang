package products

import (
	"WeMarketOnGolang/internal/models"
	"gorm.io/gorm"
)

// ProductService предоставляет логику для работы с продуктами.
type ProductService struct {
	DB *gorm.DB
}

// NewProductService создает новый экземпляр ProductService.
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

// CreateProduct создает новый продукт.
func (s *ProductService) CreateProduct(product *models.Product) error {
	if err := s.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID возвращает продукт по ID.
func (s *ProductService) GetProductByID(id int32) (*models.Product, error) {
	var product models.Product
	if err := s.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAllProducts возвращает все продукты.
func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	if err := s.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// UpdateProduct обновляет данные продукта.
func (s *ProductService) UpdateProduct(id int32, product *models.Product) error {
	if err := s.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

// DeleteProduct удаляет продукт по ID.
func (s *ProductService) DeleteProduct(id int32) error {
	if err := s.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
