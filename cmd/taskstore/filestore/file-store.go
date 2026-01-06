package filestore

import (
	"fmt"

	"github.com/task-manager-cli/cmd/taskfilemanager"
	"github.com/task-manager-cli/cmd/taskmanager"
	"github.com/task-manager-cli/cmd/utils"
)

type TaskStore struct {
	// internal storage, e.g., a map or slice
	tasks       map[int]*taskmanager.Task
	hasReadFile bool
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[int]*taskmanager.Task),
	}
}

// insert task into store
func (ts *TaskStore) Insert(task *taskmanager.Task) error {
	if err := task.Validate(); err != nil {
		return err
	}
	
	ts.tasks[task.ID] = task
	return nil
}

// retrieve task from store by ID
func (ts *TaskStore) Retrieve(id int) (*taskmanager.Task, error) {
	if task, exists := ts.tasks[id]; exists {
		return task, nil
	}

	if !ts.hasReadFile {
		// load tasks from file
		data, err := taskfilemanager.ReadTasksFromFile()
		if err != nil {
			return nil, err
		}
		// update the cache
		store, err := utils.BytesToTaskMap(data)
		if err != nil {
			return nil, err
		}

		for id, task := range store {
			ts.tasks[id] = task
		}

		ts.hasReadFile = true
		if task, exists := ts.tasks[id]; exists {
			return task, nil
		}
	}

	return nil, fmt.Errorf("task not found")
}

// update task in store
func (ts *TaskStore) Update(task *taskmanager.Task) error {
	if _, exists := ts.tasks[task.ID]; exists {
		ts.tasks[task.ID] = task
		return nil
	}
	return fmt.Errorf("task not found")
}

// delete task from store by ID
func (ts *TaskStore) Delete(id int) error {
	if _, exists := ts.tasks[id]; exists {
		delete(ts.tasks, id)
		return nil
	}
	return fmt.Errorf("task not found")
}

// list all tasks in store
func (ts *TaskStore) List() ([]*taskmanager.Task, error) {
	var taskList []*taskmanager.Task

	if !ts.hasReadFile {
		// load tasks from file
		data, err := taskfilemanager.ReadTasksFromFile()
		if err != nil {
			return nil, err
		}
		store, err := utils.BytesToTaskMap(data)
		if err != nil {
			return nil, err
		}
		// update the cache
		for id, task := range store {
			ts.tasks[id] = task
		}
		ts.hasReadFile = true
	}

	for _, task := range ts.tasks {
		taskList = append(taskList, task)
	}
	return taskList, nil
}
