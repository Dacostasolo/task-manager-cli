package filestore

import (
	"encoding/json"

	"github.com/task-manager-cli/cmd/taskfilemanager"
	"github.com/task-manager-cli/cmd/taskmanager"
)

func persistTasksToFile(ts *TaskStore) error {
	data, err := taskMapToBytes(ts.tasks)
	if err != nil {
		return err
	}

	if err := taskfilemanager.WriteTasksToFile(data); err != nil {
		return err
	}

	return nil
}

func bytesToTaskMap(data []byte) (map[int]*taskmanager.Task, error) {
	var tasks map[int]*taskmanager.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	if tasks == nil {
		tasks = make(map[int]*taskmanager.Task)
	}
	return tasks, nil
}

func taskMapToBytes(tasks map[int]*taskmanager.Task) ([]byte, error) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
