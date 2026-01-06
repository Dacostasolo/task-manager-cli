package taskstore

import taskmanager "github.com/task-manager-cli/cmd/taskmanager"

// setup file for task storage operations
type Store interface {
	// method signatures for task storage operations
	Insert(task *taskmanager.Task) error
	Retrieve(id int) (*taskmanager.Task, error)
	Update(task *taskmanager.Task) error
	Delete(id int) error
	List() ([]*taskmanager.Task, error)
}
