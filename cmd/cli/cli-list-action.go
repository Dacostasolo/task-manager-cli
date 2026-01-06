package cli

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
)

func (cli *CLI) listAction() {
	flags := parseSubcommandFlags("list")
	fmt.Println("Listing all tasks...")
	tasks, err := cli.store.List(context.Background())
	if err != nil {
		fmt.Printf("Error listing tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tTitle\tStatus\tCreated At")
	fmt.Fprintln(w, "---\t-----\t------\t----------")
	for _, task := range tasks {
		if flags.Filter != "" && task.Status != flags.Filter {
			continue
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Title, task.Status, task.CreationTimestamp.Format("2006-01-02 15:04:05"))
	}
	w.Flush()
}
