package services

import (
	"WeMarketOnGolang/internal/dto"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	"WeMarketOnGolang/internal/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JWTAuthService struct {
	JWTSecretKey []byte
}

// Генерация JWT токена
func (s *JWTAuthService) GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", user.ID),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.JWTSecretKey)
}

// Проверка пароля
func (s *JWTAuthService) ValidatePassword(storedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil
}

// Авторизация
func (s *JWTAuthService) Login(db *gorm.DB, email, password string) (*dto.LoginResponse, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
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

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
