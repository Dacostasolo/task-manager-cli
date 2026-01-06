package cli

import (
	"context"
	"fmt"

	"github.com/task-manager-cli/cmd/taskmanager"
)

func (cli *CLI) addAction() {
	commandFlags := parseSubcommandFlags("add")

	task := taskmanager.NewTask(commandFlags.Title, commandFlags.Description, commandFlags.Status)

	if err := cli.store.Insert(context.Background(), task); err != nil {
		fmt.Printf("Error adding task: %v\n", err)
		return
	}

	fmt.Printf("Task added successfully with ID: %d\n", task.ID)
}
