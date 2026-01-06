package cli

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/task-manager-cli/cmd/taskfilemanager"
	"github.com/task-manager-cli/cmd/taskstore/filestore"
)

func TestAddCommand(t *testing.T) {
	// Create a temporary file for the test
	tmpfile, err := ioutil.TempFile("", "test_tasks.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	taskfilemanager.SetTaskFilePath(tmpfile.Name())

	store := filestore.NewTaskStore()
	cli := InitCLI(store)

	// Simulate the add command
	os.Args = []string{"task-cli", "add", "-t", "Test Task", "-d", "This is a test task", "-s", "0"}
	cli.RunCommand()

	// Verify that the task was added
	tasks, err := store.List(nil)
	if err != nil {
		t.Errorf("Error getting tasks: %v", err)
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, but got %d", len(tasks))
	}

	task := tasks[0]
	if task.Title != "Test Task" {
		t.Errorf("Expected task title to be 'Test Task', but got '%s'", task.Title)
	}
	if task.Description != "This is a test task" {
		t.Errorf("Expected task description to be 'This is a test task', but got '%s'", task.Description)
	}
	if task.Status != "todo" {
		t.Errorf("Expected task status to be 'todo', but got '%s'", task.Status)
	}
}
