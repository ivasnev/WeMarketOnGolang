package handlers

import (
	"WeMarketOnGolang/internal/services/tasks"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	service *tasks.TaskService
}

func NewTaskHandler(service *tasks.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTaskInf создает новую задачу
func (h *TaskHandler) CreateTaskInf(c *gin.Context) {
	task, err := h.service.CreateTask("Inf task", h.service.InfiniteOperation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// CreateTaskClassic создает новую задачу
func (h *TaskHandler) CreateTaskClassic(c *gin.Context) {
	task, err := h.service.CreateTask("Classic task", h.service.ClassicOperation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks возвращает список всех задач
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	currentTasks := h.service.GetAllTasks()
	c.JSON(http.StatusOK, currentTasks)
}

// GetTask возвращает задачу по ID
func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CancelTask удаляет задачу по ID
func (h *TaskHandler) CancelTask(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.CancelTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
