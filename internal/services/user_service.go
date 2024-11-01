package services

import (
	"WeMarketOnGolang/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// NewUserService создаёт новый экземпляр UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// CreateUser создаёт нового пользователя
func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if err := s.DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID получает пользователя по ID
func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser обновляет данные пользователя
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	if err := s.DB.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}
