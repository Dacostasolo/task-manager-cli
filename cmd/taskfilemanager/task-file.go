package taskfilemanager

import (
	"fmt"
	"os"
	"path/filepath"
)


func getTaskFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(exePath), "tasks.json"), nil
}


func ensureTaskFileExists() error {
	path, err := getTaskFilePath()
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// Create the file if it doesn't exist
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()	
	}

	return nil
}

func ReadTasksFromFile() ([]byte, error) {
	path, err := getTaskFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get task file path: %v", err)
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
	return ensureTaskFileExists()
}