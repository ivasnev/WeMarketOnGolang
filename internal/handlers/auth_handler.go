package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Структура для работы с сервисом авторизации
type AuthHandler struct {
	authService *services.JWTAuthService
}

func NewAuthHandler(authService *services.JWTAuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Хендлер для логина с сессией
func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем IP-адрес клиента
	ipAddress := c.ClientIP()

	// Вызов метода авторизации
	loginResponse, err := h.authService.Login(request.Email, request.Password, ipAddress)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем токен в ответе
	c.JSON(http.StatusOK, loginResponse)
}

// Хендлер для логаута
func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("userID") // Предполагаем, что userID хранится в контексте после авторизации
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err := h.authService.Logout(userID.(int32))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
