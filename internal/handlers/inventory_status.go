package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services/inventoryStatus"
	"WeMarketOnGolang/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// InventoryStatusHandler представляет обработчики для инвентарных статусов.
type InventoryStatusHandler struct {
	Service *inventoryStatus.InventoryStatusService
}

// NewInventoryStatusHandler создает новый экземпляр InventoryStatusHandler.
func NewInventoryStatusHandler(service *inventoryStatus.InventoryStatusService) *InventoryStatusHandler {
	return &InventoryStatusHandler{Service: service}
}

// CreateInventoryStatus создает новый инвентарный статус.
// @Summary Создать инвентарный статус
// @Description Создает новый инвентарный статус.
// @Tags v1/inventory_status
// @Accept json
// @Produce json
// @Param status body dto.InventoryStatusDTO true "Информация о статусе"
// @Success 201 {object} map[string]interface{} "Созданный статус"
// @Failure 400 {object} dto.ErrorResponse "Некорректные данные"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/inventory_status [post]
func (h *InventoryStatusHandler) CreateInventoryStatus(c *gin.Context) {
	var dto dto.InventoryStatusDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	status := &models.InventoryStatus{Name: dto.Name}
	if err := h.Service.CreateInventoryStatus(status); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": status.ID, "name": status.Name})
}

// GetInventoryStatusByID возвращает инвентарный статус по ID.
// @Summary Получить инвентарный статус
// @Description Возвращает инвентарный статус по ID.
// @Tags v1/inventory_status
// @Produce json
// @Param id path int true "ID статуса"
// @Success 200 {object} map[string]interface{} "Информация о статусе"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID"
// @Failure 404 {object} map[string]interface{} "Статус не найден"
// @Security BearerAuth
// @Router /v1/inventory_status/{id} [get]
func (h *InventoryStatusHandler) GetInventoryStatusByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	status, err := h.Service.GetInventoryStatusByID(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": status.ID, "name": status.Name})
}

// GetAllInventoryStatuses возвращает все инвентарные статусы.
// @Summary Получить все инвентарные статусы
// @Description Возвращает список всех инвентарных статусов.
// @Tags v1/inventory_status
// @Produce json
// @Success 200 {array} models.InventoryStatus "Список статусов"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/inventory_status [get]
func (h *InventoryStatusHandler) GetAllInventoryStatuses(c *gin.Context) {
	statuses, err := h.Service.GetAllInventoryStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statuses"})
		return
	}

	c.JSON(http.StatusOK, statuses)
}

// UpdateInventoryStatus обновляет инвентарный статус по ID.
// @Summary Обновить инвентарный статус
// @Description Обновляет данные существующего инвентарного статуса.
// @Tags v1/inventory_status
// @Accept json
// @Produce json
// @Param id path int true "ID статуса"
// @Param status body dto.InventoryStatusDTO true "Информация для обновления"
// @Success 200 {object} map[string]interface{} "Обновленный статус"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID или данные"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/inventory_status/{id} [put]
func (h *InventoryStatusHandler) UpdateInventoryStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var dto dto.InventoryStatusDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	status := &models.InventoryStatus{ID: int32(id), Name: dto.Name}
	if err := h.Service.UpdateInventoryStatus(int32(id), status); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": status.ID, "name": status.Name})
}

// DeleteInventoryStatus удаляет инвентарный статус по ID.
// @Summary Удалить инвентарный статус
// @Description Удаляет существующий инвентарный статус по ID.
// @Tags v1/inventory_status
// @Produce json
// @Param id path int true "ID статуса"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} dto.ErrorResponse "Некорректный ID"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/inventory_status/{id} [delete]
func (h *InventoryStatusHandler) DeleteInventoryStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteInventoryStatus(int32(id)); err != nil {
		statusCode, response := utils.HandleDBError(err)
		c.JSON(statusCode, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted"})
}
