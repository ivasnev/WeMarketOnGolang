package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services"
	"WeMarketOnGolang/internal/utils"
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

	userIDInt, err := strconv.Atoi(userID.(string)) // преобразуем строку в int
	if err != nil {
		return 0, errors.New("Invalid userID")
	}
	return int32(userIDInt), nil
}

// Register регистрирует нового пользователя
// @Summary Регистрация пользователя
// @Description Создает нового пользователя с указанными данными
// @Tags v1/auth
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDTO true "Данные пользователя"
// @Success 201 {object} map[string]interface{} "Успешная регистрация"
// @Failure 400 {object} map[string]interface{} "Ошибка ввода"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Router /v1/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var request dto.CreateUserDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		RoleID:       1,
	}

	if err := h.userService.CreateUser(&user); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// GetCurrentUser возвращает текущего пользователя
// @Summary Получить текущего пользователя
// @Description Возвращает данные текущего авторизованного пользователя
// @Tags v1/users
// @Produce json
// @Success 200 {object} dto.UserResponseDTO "Данные пользователя"
// @Failure 401 {object} map[string]interface{} "Пользователь не авторизован"
// @Security BearerAuth
// @Router /v1/users/me [get]
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

// UpdateCurrentUser обновляет данные текущего пользователя
// @Summary Обновить текущего пользователя
// @Description Обновляет данные текущего авторизованного пользователя
// @Tags v1/users
// @Accept json
// @Produce json
// @Param user body dto.UpdateUserDTO true "Обновляемые данные пользователя"
// @Success 200 {object} map[string]interface{} "Успешное обновление"
// @Failure 400 {object} map[string]interface{} "Ошибка ввода"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/users/me [patch]
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

// GetUserByID возвращает пользователя по ID
// @Summary Получить пользователя по ID
// @Description Возвращает данные пользователя на основе идентификатора
// @Tags v1/users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} dto.UserResponseDTO "Данные пользователя"
// @Failure 400 {object} map[string]interface{} "Некорректный ID"
// @Failure 404 {object} map[string]interface{} "Пользователь не найден"
// @Security BearerAuth
// @Router /v1/users/{id} [get]
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
