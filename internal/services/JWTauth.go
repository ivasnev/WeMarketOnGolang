package services

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type JWTAuthService struct {
	jwtSecretKey []byte
	db           *gorm.DB
}

func NewJWTAuthService(jwtSecretKey []byte, db *gorm.DB) *JWTAuthService {
	return &JWTAuthService{jwtSecretKey: jwtSecretKey, db: db}
}

// Генерация JWT токена
func (s *JWTAuthService) GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", user.ID),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecretKey)
}

// Проверка пароля
func (s *JWTAuthService) ValidatePassword(storedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil
}

// Авторизация с созданием сессии
func (s *JWTAuthService) Login(email, password, ipAddress string) (*dto.LoginResponse, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !s.ValidatePassword(user.PasswordHash, password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.GenerateJWT(&user)
	if err != nil {
		return nil, err
	}

	// Создаем запись о сессии
	startDate := time.Now()
	session := models.UserSession{
		UserID:    user.ID,
		StartDate: &startDate,
		IPAddress: &ipAddress,
	}

	if err := s.db.Create(&session).Error; err != nil {
		return nil, errors.New("failed to create session")
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

// Logout Логаут с записью времени завершения сессии
func (s *JWTAuthService) Logout(userID int32) error {
	var session models.UserSession
	err := s.db.Where("user_id = ? AND end_date IS NULL", userID).First(&session).Error
	if err != nil {
		return errors.New("no active session found")
	}

	endDate := time.Now()
	session.EndDate = &endDate

	// Обновляем запись о сессии
	if err := s.db.Save(&session).Error; err != nil {
		return errors.New("failed to update session end time")
	}

	return nil
}
