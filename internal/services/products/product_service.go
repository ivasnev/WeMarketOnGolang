package products

import (
	"WeMarketOnGolang/internal/dto"
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
func (s *ProductService) GetAllProducts(filter *dto.ProductFilter) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	query := s.DB.Model(&models.Product{})

	// Применяем фильтры
	if filter.Name != nil {
		query = query.Where("name ILIKE ?", "%"+*filter.Name+"%") // Для поиска по подстроке
	}
	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}

	// Считаем общее количество записей для пагинации
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Пагинация
	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		query = query.Limit(filter.PageSize).Offset(offset)
	}

	// Выполняем запрос
	if err := query.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
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
