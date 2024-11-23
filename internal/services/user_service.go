package services

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Хэшируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)

	// Добавляем пользователя в базу
	result := s.DB.Create(user)
	return result.Error
}

func (s *UserService) GetUserByID(id int32) (*models.User, error) {
	var user models.User
	result := s.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) UpdateUser(id int32, user *models.User) error {
	var existingUser models.User
	result := s.DB.First(&existingUser, id)
	if result.Error != nil {
		return result.Error
	}

	// Обновляем данные пользователя
	result = s.DB.Model(&existingUser).Updates(user)
	return result.Error
}

func (s *UserService) DeleteUser(id int32) error {
	result := s.DB.Delete(&models.User{}, id)
	return result.Error
}

func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	result := s.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	// Проверяем пароль
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

[]byte("your_secret_key")
