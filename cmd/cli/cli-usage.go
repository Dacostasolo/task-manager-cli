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
}

func cliUsage() {
	usageText := `Task Manager CLI - A simple task management application
Usage:
taskmanager [command] [options]
Commands:
add       Add a new task
list      List all tasks
update    Update an existing task
delete    Delete a task
get       Get details of a specific task
save      Save tasks to persistent storage

Options:
-t, -title         Title of the task (required for add)
-d, -description   Description of the task (optional for add)
-s, -status        Status of the task (0 = todo, 1 = in-progress, 2 = done) (optional for add/update)
-id                ID of the task (required for update, delete, get)

Examples:
taskmanager add -t "Buy groceries" -d "Milk, Bread, Eggs" -s "Pending"
taskmanager list
taskmanager update -id 123456 -s "Completed"
taskmanager delete -id 123456
taskmanager get -id 123456
taskmanager save
`
	fmt.Println(usageText)
}

func parseSubcommandFlags(action string) *CommandFlags {
	flagAction := flag.NewFlagSet(action, flag.ExitOnError)
	id := flagAction.Int("id", 0, "ID of the task")
	title := flagAction.String("t", "", "Title of the task")
	description := flagAction.String("d", "", "Description of the task")
	status := flagAction.Int("s", 0, "Status of the task")

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
	}
}
