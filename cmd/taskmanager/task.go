package taskmanager

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	StatusPending   = "Pending"
	StatusInProgress = "In Progress"
	StatusCompleted  = "Completed"
)

type Task struct {
	ID          int `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof='Pending' 'In Progress' 'Completed'"`
}


func (t *Task) Validate() error {
	var validate = validator.New()
	return validate.Struct(t) 
}


func NewTask( title, description, status string) *Task {

	if status == "" {
		status = StatusPending
	}
	
	return &Task{
		ID:          time.Now().Nanosecond(), // Simple unique ID based on timestamp
		Title:       title,
		Description: description,
		Status:      status,
	}
}

func (t *Task) UpdateStatus(newStatus string) {
	t.Status = newStatus
}

func (t *Task) String() string {
	return fmt.Sprintf("Task(ID: %d, Title: %s, Description: %s, Status: %s)", t.ID, t.Title, t.Description, t.Status)
}