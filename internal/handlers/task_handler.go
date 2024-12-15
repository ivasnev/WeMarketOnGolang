package handlers

import (
	"WeMarketOnGolang/internal/dto"
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

// CreateTaskInf создает задачу с бесконечной операцией
// @Summary Создать задачу (бесконечная операция)
// @Description Создает новую задачу, выполняющую бесконечную операцию
// @Tags v1/tasks
// @Accept json
// @Produce json
// @Success 201 {object} string "Созданная задача"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/tasks/inf [post]
func (h *TaskHandler) CreateTaskInf(c *gin.Context) {
	task, err := h.service.CreateTask("Inf task", h.service.InfiniteOperation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// CreateTaskClassic создает задачу с классической операцией
// @Summary Создать задачу (классическая операция)
// @Description Создает новую задачу, выполняющую классическую операцию
// @Tags v1/tasks
// @Accept json
// @Produce json
// @Success 201 {object} string "Созданная задача"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/tasks/classic [post]
func (h *TaskHandler) CreateTaskClassic(c *gin.Context) {
	task, err := h.service.CreateTask("Classic task", h.service.ClassicOperation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetAllTasks возвращает список всех задач
// @Summary Получить все задачи
// @Description Возвращает список всех активных задач
// @Tags v1/tasks
// @Accept json
// @Produce json
// @Success 200 {array} []dto.TaskResponse "Список задач"
// @Security BearerAuth
// @Router /v1/tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	currentTasks := h.service.GetAllTasks()
	c.JSON(http.StatusOK, currentTasks)
}

// GetTask возвращает задачу по ID
// @Summary Получить задачу по ID
// @Description Возвращает задачу на основе переданного идентификатора
// @Tags v1/tasks
// @Accept json
// @Produce json
// @Param id path string true "ID задачи"
// @Success 200 {object} dto.TaskResponse "Данные задачи"
// @Failure 404 {object} map[string]interface{} "Задача не найдена"
// @Security BearerAuth
// @Router /v1/tasks/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, dto.TaskResponse{
		ID:          task.ID,
		Status:      task.Status,
		Error:       task.Error,
		Progress:    task.Progress,
		Description: task.Description,
	})
}

// CancelTask удаляет задачу по ID
// @Summary Удалить задачу
// @Description Удаляет задачу по ее идентификатору
// @Tags v1/tasks
// @Accept json
// @Produce json
// @Param id path string true "ID задачи"
// @Success 200 {object} map[string]interface{} "Задача успешно удалена"
// @Failure 500 {object} dto.ErrorResponse "Ошибка сервера"
// @Security BearerAuth
// @Router /v1/tasks/{id} [delete]
func (h *TaskHandler) CancelTask(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.CancelTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
