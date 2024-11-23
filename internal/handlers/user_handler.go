package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) getUserIdFromContext(c *gin.Context) (int32, error) {
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

func (h *UserHandler) Register(c *gin.Context) {
	var request dto.CreateUserDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хешируем пароль перед сохранением
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:        request.Email,
		PasswordHash: string(passwordHash),
		Name:         request.Name,
		Phone:        request.Phone,
		Address:      request.Address,
		RoleID:       1, // Роль по умолчанию
	}

	if err := h.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, err := h.getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	response := dto.UserResponseDTO{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
		RoleID:  user.RoleID,
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}

func (h *UserHandler) UpdateCurrentUser(c *gin.Context) {
	userID, err := h.getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var request dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Обновление только тех полей, которые переданы в запросе
	if request.Name != nil {
		user.Name = *request.Name
	}
	if request.Phone != nil {
		user.Phone = request.Phone
	}
	if request.Address != nil {
		user.Address = request.Address
	}

	if err := h.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.userService.GetUserByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	response := dto.UserResponseDTO{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
		RoleID:  user.RoleID,
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}
