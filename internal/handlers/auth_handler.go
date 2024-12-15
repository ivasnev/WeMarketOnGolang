package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Структура для работы с сервисом авторизации
type AuthHandler struct {
	authService *services.JWTAuthService
}

func NewAuthHandler(authService *services.JWTAuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login
// @Summary Авторизация пользователя
// @Description Принимает логин и пароль, возвращает JWT токен
// @Tags v1/auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Учетные данные пользователя"
// @Success 200 {object} dto.LoginResponse "JWT токен"
// @Failure 400 {object} map[string]interface{} "Неверные данные"
// @Failure 401 {object} map[string]interface{} "Ошибка авторизации"
// @Router /v1/auth/jwt/login [post]
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

func (h *AuthHandler) getUserIdFromContext(c *gin.Context) (int32, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("No userID found in context")
	}

	// Преобразуем userID в int32
	userIDInt, err := strconv.Atoi(userID.(string)) // преобразуем строку в int
	if err != nil {
		return 0, errors.New("Invalid userID")
	}
	return int32(userIDInt), nil
}

// Logout
// @Summary Выход пользователя
// @Description Завершение сессии пользователя
// @Tags v1/auth
// @Produce json
// @Success 200 {object} map[string]interface{} "Успешный выход"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/auth/jwt/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	userID, err := h.getUserIdFromContext(c) // Предполагаем, что userID хранится в контексте после авторизации
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err = h.authService.Logout(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
