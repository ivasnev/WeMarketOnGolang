package categories

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

// CreateCategory создает новую категорию в базе данных
func (s *CategoryService) CreateCategory(req dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:             req.Name,
		Description:      req.Description,
		ParentCategoryID: req.ParentCategoryID,
	}

	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:               category.ID,
		Name:             category.Name,
		Description:      category.Description,
		ParentCategoryID: category.ParentCategoryID,
	}, nil
}

// GetCategoryByID возвращает категорию по ID
func (s *CategoryService) GetCategoryByID(id int) (*dto.CategoryResponse, error) {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	response := &dto.CategoryResponse{
		ID:               category.ID,
		Name:             category.Name,
		Description:      category.Description,
		ParentCategoryID: category.ParentCategoryID,
	}

	return response, nil
}

// GetAllCategories возвращает список всех категорий
func (s *CategoryService) GetAllCategories() ([]dto.CategoryResponse, error) {
	var categories []models.Category
	if err := s.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = dto.CategoryResponse{
			ID:               category.ID,
			Name:             category.Name,
			Description:      category.Description,
			ParentCategoryID: category.ParentCategoryID,
		}
	}

	return responses, nil
}

// UpdateCategory обновляет категорию в базе данных
func (s *CategoryService) UpdateCategory(id int, req dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	if req.Name != nil {
		category.Name = *req.Name
	}
	if req.Description != nil {
		category.Description = req.Description
	}
	if req.ParentCategoryID != nil {
		category.ParentCategoryID = req.ParentCategoryID
	}

	if err := s.db.Save(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:               category.ID,
		Name:             category.Name,
		Description:      category.Description,
		ParentCategoryID: category.ParentCategoryID,
	}, nil
}

// DeleteCategory удаляет категорию по ID
func (s *CategoryService) DeleteCategory(id int) error {
	if err := s.db.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
