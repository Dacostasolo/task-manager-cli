package cli

import (
	"flag"
	"fmt"
	"os"
)

type CommandFlags struct {
	Title       string
	Description string
	Status      string
	Persist     bool
	ID          int
	Filter      string
}

func cliUsage() {
	usageText := `
🚀 Welcome to Task Tracker CLI! 🚀

A simple yet powerful command-line tool for managing your tasks.

Usage:
  task-cli [command] [options]

Commands:
  add       డ్ Adds a new task
  list      📜 Lists all tasks
  update    ♻️  Updates an existing task
  delete    ❌ Deletes a task
  get       ℹ️  Gets details of a specific task
  save      💾 Saves tasks to persistent storage

Options:
  -t, -title         📝 Title of the task (required for 'add')
  -d, -description   📋 Description of the task (optional for 'add')
  -s, -status        📊 Status of the task (0 = todo, 1 = in-progress, 2 = done) (optional for 'add'/'update')
  -id                🆔 ID of the task (required for 'update', 'delete', 'get')
  -f, -filter        🔎 Filter tasks by status when listing (optional for 'list')

Examples:
  # Add a new task
  task-cli add -t "Buy groceries" -d "Milk, Bread, Eggs"

  # List all tasks
  task-cli list

  # Update a task's status to 'in-progress'
  task-cli update -id 123456 -s 1

  # Delete a task
  task-cli delete -id 123456

  # Get details of a task
  task-cli get -id 123456

  # Save all tasks
  task-cli save
`
	fmt.Println(usageText)
}

func parseSubcommandFlags(action string) *CommandFlags {
	flagAction := flag.NewFlagSet(action, flag.ExitOnError)
	id := flagAction.Int("id", 0, "ID of the task")
	title := flagAction.String("t", "", "Title of the task")
	description := flagAction.String("d", "", "Description of the task")
	status := flagAction.Int("s", 0, "Status of the task")
	filter := flagAction.String("f", "", "Filter tasks by status")

	flagAction.Parse(os.Args[2:])

	stringStatus := ""
	switch *status {
	case 0:
		stringStatus = "todo"
	case 1:
		stringStatus = "in-progress"
	case 2:
		stringStatus = "done"
	default:
		stringStatus = "todo"
	}

	return &CommandFlags{
		Title:       *title,
		Description: *description,
		Status:      stringStatus,
		ID:          *id,
		Filter:      *filter,
	}
}
