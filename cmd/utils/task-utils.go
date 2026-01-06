package utils

import (
	"encoding/json"

	taskmanager "github.com/task-manager-cli/cmd/taskmanager"
)

// coverstion from bytes to map[int]Task
func BytesToTaskMap(data []byte) (map[int]*taskmanager.Task, error) {
	var tasks map[int]*taskmanager.Task
	err := json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// conversion from map[int]Task to bytes
func TaskMapToBytes(tasks map[int]*taskmanager.Task) ([]byte, error) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}