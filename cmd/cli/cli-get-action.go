package cli

import (
	"context"
	"fmt"
)

func (cli *CLI) getAction() {
	commandFlags := parseSubcommandFlags("get")

	if commandFlags.ID == 0 {
		fmt.Println("Error: Task ID is required for retrieval.")
		return
	}

	task, err := cli.store.Retrieve(context.Background(), commandFlags.ID)
	if err != nil {
		fmt.Printf("Error retrieving task: %v\n", err)
		return
	}

	fmt.Println("--- Task Details ---")
	fmt.Printf("ID:          %d\n", task.ID)
	fmt.Printf("Title:       %s\n", task.Title)
	fmt.Printf("Description: %s\n", task.Description)
	fmt.Printf("Status:      %s\n", task.Status)
	fmt.Printf("Created At:  %s\n", task.CreationTimestamp.Format("2006-01-02 15:04:05"))
}