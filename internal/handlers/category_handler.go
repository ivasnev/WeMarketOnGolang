package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/services/categories"
	"WeMarketOnGolang/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CategoryHandler обрабатывает запросы, связанные с категориями.
type CategoryHandler struct {
	service *categories.CategoryService
}

// NewCategoryHandler создает новый CategoryHandler.
func NewCategoryHandler(service *categories.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// CreateCategory создает новую категорию.
// @Summary Создание категории
// @Description Создает новую категорию на основе предоставленных данных.
// @Tags v1/categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Данные для создания категории"
// @Success 201 {object} dto.CategoryResponse "Созданная категория"
// @Failure 400 {object} dto.ErrorResponse "Ошибка валидации входных данных"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/category [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.CreateCategory(req)
	if err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategory возвращает категорию по ID.
// @Summary Получение категории по ID
// @Description Возвращает категорию по её идентификатору.
// @Tags v1/categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} dto.CategoryResponse "Найденная категория"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID категории"
// @Failure 404 {object} map[string]interface{} "Категория не найдена"
// @Security BearerAuth
// @Router /v1/category/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetAllCategories возвращает список всех категорий.
// @Summary Получение списка категорий
// @Description Возвращает список всех категорий.
// @Tags v1/categories
// @Produce json
// @Success 200 {array} dto.CategoryResponse "Список категорий"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/category [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categoriesList, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	c.JSON(http.StatusOK, categoriesList)
}

// UpdateCategory обновляет информацию о категории.
// @Summary Обновление категории
// @Description Обновляет данные существующей категории.
// @Tags v1/categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Param category body dto.UpdateCategoryRequest true "Данные для обновления категории"
// @Success 200 {object} dto.CategoryResponse "Обновленная категория"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID категории или ошибка валидации"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/category/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.service.UpdateCategory(id, req)
	if err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory удаляет категорию по ID.
// @Summary Удаление категории
// @Description Удаляет категорию по её идентификатору.
// @Tags v1/categories
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID категории"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/category/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.service.DeleteCategory(id); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
