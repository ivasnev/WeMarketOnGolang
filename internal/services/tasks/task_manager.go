package tasks

import (
	"WeMarketOnGolang/internal/dto"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// TaskService отвечает за управление задачами
type TaskService struct {
	tasks map[string]*dto.Task
	mu    sync.Mutex
	sema  chan struct{} // Семафор для ограничения задач
	limit int           // Лимит на количество задач
}

// NewTaskService создает новый экземпляр TaskService
func NewTaskService(limit int) *TaskService {
	return &TaskService{
		tasks: make(map[string]*dto.Task),
		sema:  make(chan struct{}, limit), // Буферизованный канал для семафора
		limit: limit,
	}
}

// CreateTask создает новую задачу и выполняет переданную функцию с возможностью отмены
func (s *TaskService) CreateTask(info string, fn func(ctx context.Context) error) (string, error) {
	// Проверяем семафор перед созданием задачи
	select {
	case s.sema <- struct{}{}: // Если есть место, продолжаем
	default:
		// Если все места заняты, возвращаем ошибку
		return "", errors.New("task limit reached, please try again later")
	}

	// Создаем новую задачу
	taskID := fmt.Sprintf("%d", time.Now().UnixNano())
	task := &dto.Task{
		ID:          taskID,
		Status:      "pending",
		Progress:    0,
		Description: info,
	}

	// Добавляем задачу в список
	s.mu.Lock()
	s.tasks[taskID] = task
	s.mu.Unlock()

	// Запуск задачи в отдельной горутине
	go func() {
		s.mu.Lock()
		task.Status = dto.StatusRunning
		s.mu.Unlock()

		ctx, cancel := context.WithCancel(context.Background()) // Устанавливаем контекст
		task.Cancel = cancel

		errChan := make(chan error)
		go func() {
			errChan <- fn(ctx)
		}()
		err := <-errChan
		s.mu.Lock()
		if err != nil {
			task.Status = dto.StatusFailed
			task.Progress = 0
			task.Description += fmt.Sprintf(" - Error: %v", err)
		} else {
			task.Status = dto.StatusSuccess
			task.Progress = 100
		}
		s.mu.Unlock()

		// Освобождаем место в семафоре
		<-s.sema
	}()

	return taskID, nil
}

// Пример бесконечной операции
func (s *TaskService) InfiniteOperation(ctx context.Context) error {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done(): // Проверяем отмену
			return errors.New("operation canceled")
		default:
			fmt.Println("Работаю бесконечно...")
		}
	}
}

// Пример операции с поддержкой отмены
func (s *TaskService) ClassicOperation(ctx context.Context) error {

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done(): // Проверяем отмену
			return errors.New("operation canceled")
		default:
			fmt.Printf("Выполняю итерацию %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}

	return nil
}

func (s *TaskService) CancelTask(taskID string) error {

	if task, exists := s.tasks[taskID]; exists {
		if task.Status == dto.StatusRunning {
			task.Cancel()
			task.UpdateTaskStatus(dto.StatusCanceled)
			return nil
		}
		return errors.New("task not running")
	}

	return errors.New("task not found")
}

// GetTaskByID возвращает задачу по ID
func (s *TaskService) GetTaskByID(id string) (*dto.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}

	return task, nil
}

// GetAllTasks возвращает список всех задач
func (s *TaskService) GetAllTasks() []*dto.TaskResponse {
	s.mu.Lock()
	defer s.mu.Unlock()

	var taskList []*dto.TaskResponse
	for _, task := range s.tasks {
		taskList = append(taskList, &dto.TaskResponse{
			ID:          task.ID,
			Status:      task.Status,
			Error:       task.Error,
			Progress:    task.Progress,
			Description: task.Description,
		})
	}

	return taskList
}

// DeleteTask удаляет задачу по ID
func (s *TaskService) DeleteTask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[id]
	if !exists {
		return fmt.Errorf("task not found")
	}

	delete(s.tasks, id)
	return nil
}
