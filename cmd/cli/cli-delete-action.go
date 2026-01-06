package cli

import (
	"context"
	"fmt"
)

func (cli *CLI) deleteAction() {
	commandFlags := parseSubcommandFlags("delete")

	if commandFlags.ID == 0 {
		fmt.Println("Error: Task ID is required for deletion.")
		return
	}

	if err := cli.store.Delete(context.Background(), commandFlags.ID); err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		return
	}

	fmt.Println("Task deleted successfully with ID:", commandFlags.ID)
}