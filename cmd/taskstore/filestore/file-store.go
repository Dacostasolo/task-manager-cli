package filestore

import (
	"context"
	"fmt"
	"sync"

	"github.com/task-manager-cli/cmd/taskfilemanager"
	"github.com/task-manager-cli/cmd/taskmanager"
)

type TaskStore struct {
	mu          sync.RWMutex
	tasks       map[int]*taskmanager.Task
	hasReadFile bool
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[int]*taskmanager.Task),
	}
}

func (ts *TaskStore) Insert(ctx context.Context, task *taskmanager.Task) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if err := ts.loadTasksFromFile(); err != nil {
		return err
	}

	maxID := 0
	for _, t := range ts.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task.ID = maxID + 1

	if err := task.Validate(); err != nil {
		return err
	}

	ts.tasks[task.ID] = task
	return persistTasksToFile(ts)
}

func (ts *TaskStore) Retrieve(ctx context.Context, id int) (*taskmanager.Task, error) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	if err := ts.loadTasksFromFile(); err != nil {
		return nil, err
	}

	if task, exists := ts.tasks[id]; exists {
		return task, nil
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (ts *TaskStore) Update(ctx context.Context, task *taskmanager.Task) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if err := task.Validate(); err != nil {
		return err
	}

	if err := ts.loadTasksFromFile(); err != nil {
		return err
	}

	if _, exists := ts.tasks[task.ID]; exists {
		ts.tasks[task.ID] = task
		return persistTasksToFile(ts)
	}

	return fmt.Errorf("task with ID %d not found", task.ID)
}

func (ts *TaskStore) Delete(ctx context.Context, id int) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if err := ts.loadTasksFromFile(); err != nil {
		return err
	}

	if _, exists := ts.tasks[id]; exists {
		delete(ts.tasks, id)
		return persistTasksToFile(ts)
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (ts *TaskStore) List(ctx context.Context) ([]*taskmanager.Task, error) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	if err := ts.loadTasksFromFile(); err != nil {
		return nil, err
	}

	var taskList []*taskmanager.Task
	for _, task := range ts.tasks {
		taskList = append(taskList, task)
	}
	return taskList, nil
}

func (ts *TaskStore) loadTasksFromFile() error {
	if ts.hasReadFile {
		return nil
	}

	data, err := taskfilemanager.ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(data) > 0 {
		store, err := bytesToTaskMap(data)
		if err != nil {
			return err
		}
		for id, task := range store {
			ts.tasks[id] = task
		}
	}

	ts.hasReadFile = true
	return nil
}

func (ts *TaskStore) Persist(ctx context.Context) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	return persistTasksToFile(ts)
}
