package cli

import (
	"context"
	"fmt"
)

func (cli *CLI) persistAction() {
	fmt.Println("Saving tasks...")

	if err := cli.store.Persist(context.Background()); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}

	fmt.Println("Tasks persisted successfully.")
}
