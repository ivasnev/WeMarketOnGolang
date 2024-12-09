package dto

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TaskStatus string

const (
	StatusPending  TaskStatus = "Pending"
	StatusRunning  TaskStatus = "Running"
	StatusSuccess  TaskStatus = "Success"
	StatusFailed   TaskStatus = "Failed"
	StatusCanceled TaskStatus = "Canceled"
)

type Task struct {
	ID          string
	Status      TaskStatus
	Error       error
	mu          sync.Mutex
	Cancel      context.CancelFunc
	Progress    int    `json:"progress"`
	Description string `json:"description"`
}

func (ts *Task) UpdateTaskStatus(status TaskStatus) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	fmt.Println(fmt.Sprintf("%s: status for %s changed %s -> %s",
		time.Now().Format("15:04:05"),
		ts.ID,
		ts.Status,
		status))

	ts.Status = status

}

func (ts *Task) GetState() {
	fmt.Printf("Task %s - \"%s\" State: %s \n", ts.ID, ts.Description, ts.Status)
}

type TaskRequest struct {
	Info string `json:"info"`
}

type TaskResponse struct {
	ID          string `json:"id"`
	Status      TaskStatus
	Error       error
	Progress    int    `json:"progress"`
	Description string `json:"description"`
}
