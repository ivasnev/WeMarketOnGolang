package handlers

import (
	"WeMarketOnGolang/internal/dto"
	"WeMarketOnGolang/internal/models"
	"WeMarketOnGolang/internal/services/inventoryStatus"
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
func (h *InventoryStatusHandler) CreateInventoryStatus(c *gin.Context) {
	var dto dto.InventoryStatusDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	status := &models.InventoryStatus{Name: dto.Name}
	if err := h.Service.CreateInventoryStatus(status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create status"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": status.ID, "name": status.Name})
}

// GetInventoryStatusByID возвращает инвентарный статус по ID.
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
func (h *InventoryStatusHandler) GetAllInventoryStatuses(c *gin.Context) {
	statuses, err := h.Service.GetAllInventoryStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch statuses"})
		return
	}

	c.JSON(http.StatusOK, statuses)
}

// UpdateInventoryStatus обновляет инвентарный статус по ID.
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": status.ID, "name": status.Name})
}

// DeleteInventoryStatus удаляет инвентарный статус по ID.
func (h *InventoryStatusHandler) DeleteInventoryStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteInventoryStatus(int32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted"})
}
