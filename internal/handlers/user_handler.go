package handlers

import (
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewUserHandler(userService *services.UserService, authService *services.AuthService) *UserHandler {
	return &UserHandler{UserService: userService, AuthService: authService}
}

// Регистрация нового пользователя
func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Вход пользователя
func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.AuthenticateUser(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := h.AuthService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Получение данных пользователя по ID
func (h *UserHandler) GetUser(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Преобразуем int в int32
	userID := int32(num)
	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Обновление данных пользователя
func (h *UserHandler) UpdateUser(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Преобразуем int в int32
	userID := int32(num)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.UpdateUser(userID, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// Удаление пользователя
func (h *UserHandler) DeleteUser(c *gin.Context) {

	num, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Преобразуем int в int32
	userID := int32(num)
	if err := h.UserService.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
