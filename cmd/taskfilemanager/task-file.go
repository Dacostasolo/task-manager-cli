package taskfilemanager

import (
	"fmt"
	"os"
	"path/filepath"
)

var taskFilePath string

func SetTaskFilePath(path string) {
	taskFilePath = path
}

func getTaskFilePath() (string, error) {
	if taskFilePath != "" {
		return taskFilePath, nil
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user config directory: %v", err)
	}

	appConfigDir := filepath.Join(configDir, "task-tracker-cli")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create app config directory: %v", err)
	}

	return filepath.Join(appConfigDir, "tasks.json"), nil
}

func ensureTaskFileExists() (string, error) {
	path, err := getTaskFilePath()
	if err != nil {
		return "", err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Create the file with an empty JSON object
		if err := os.WriteFile(path, []byte("{}"), 0644); err != nil {
			return "", err
		}
	}

	return path, nil
}

func ReadTasksFromFile() ([]byte, error) {
	path, err := ensureTaskFileExists()
	if err != nil {
		return nil, fmt.Errorf("failed to ensure task file exists: %v", err)
	}

	return os.ReadFile(path)
}

func WriteTasksToFile(data []byte) error {
	path, err := getTaskFilePath()
	if err != nil {
		return fmt.Errorf("failed to get task file path: %v", err)
	}

	return os.WriteFile(path, data, 0644)
}

func InitTaskFile() error {
	_, err := ensureTaskFileExists()
	return err
}
