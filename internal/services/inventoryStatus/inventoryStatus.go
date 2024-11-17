package inventoryStatus

import (
	"WeMarketOnGolang/internal/models"
	"gorm.io/gorm"
)

// InventoryStatusService предоставляет логику для работы с инвентарными статусами.
type InventoryStatusService struct {
	DB *gorm.DB
}

// NewInventoryStatusService создает новый экземпляр InventoryStatusService.
func NewInventoryStatusService(db *gorm.DB) *InventoryStatusService {
	return &InventoryStatusService{DB: db}
}

// CreateInventoryStatus создает новый инвентарный статус.
func (s *InventoryStatusService) CreateInventoryStatus(status *models.InventoryStatus) error {
	if err := s.DB.Create(status).Error; err != nil {
		return err
	}
	return nil
}

// GetInventoryStatusByID возвращает инвентарный статус по ID.
func (s *InventoryStatusService) GetInventoryStatusByID(id int32) (*models.InventoryStatus, error) {
	var status models.InventoryStatus
	if err := s.DB.First(&status, id).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

// GetAllInventoryStatuses возвращает все инвентарные статусы.
func (s *InventoryStatusService) GetAllInventoryStatuses() ([]models.InventoryStatus, error) {
	var statuses []models.InventoryStatus
	if err := s.DB.Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

// UpdateInventoryStatus обновляет данные инвентарного статуса по ID.
func (s *InventoryStatusService) UpdateInventoryStatus(id int32, status *models.InventoryStatus) error {
	if err := s.DB.Model(&models.InventoryStatus{}).Where("id = ?", id).Updates(status).Error; err != nil {
		return err
	}
	return nil
}

// DeleteInventoryStatus удаляет инвентарный статус по ID.
func (s *InventoryStatusService) DeleteInventoryStatus(id int32) error {
	if err := s.DB.Delete(&models.InventoryStatus{}, id).Error; err != nil {
		return err
	}
	return nil
}
