package cli

import (
	"context"
	"fmt"
)

func (cli *CLI) updateAction() {
	commandFlags := parseSubcommandFlags("update")

	if commandFlags.ID == 0 {
		fmt.Println("Error: Task ID is required for update.")
		return
	}

	task, err := cli.store.Retrieve(context.Background(), commandFlags.ID)
	if err != nil {
		fmt.Printf("Error retrieving task: %v\n", err)
		return
	}

	if commandFlags.Title != "" {
		task.Title = commandFlags.Title
	}
	if commandFlags.Description != "" {
		task.Description = commandFlags.Description
	}
	if commandFlags.Status != "" {
		task.Status = commandFlags.Status
	}

	if err := cli.store.Update(context.Background(), task); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Println("Task updated successfully:", task)
}

