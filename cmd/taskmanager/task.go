package taskmanager

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	StatusTodo  = "todo"
	StatusInProgress = "in-progress"
	StatusDone  = "done"
)

var validate = validator.New()

type Task struct {
	ID          int `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof='todo' 'in-progress' 'done'"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
}

func (t *Task) Validate() error {
	return validate.Struct(t)
}

func NewTask(title, description, status string) *Task {
	if status == "" {
		status = StatusTodo
	}

	return &Task{
		Title:             title,
		Description:       description,
		Status:            status,
		CreationTimestamp: time.Now(),
	}
}

func (t *Task) UpdateStatus(newStatus string) {
	switch newStatus {
	case StatusTodo, StatusInProgress, StatusDone:
		t.Status = newStatus
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("Task(ID: %d, Title: %s, Description: %s, Status: %s, Created At: %s)", t.ID, t.Title, t.Description, t.Status, t.CreationTimestamp.Format("2006-01-02 15:04:05"))
}