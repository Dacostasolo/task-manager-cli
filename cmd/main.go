package main

import (
	"fmt"

	"github.com/task-manager-cli/cmd/cli"
	"github.com/task-manager-cli/cmd/taskstore/filestore"
)

func main() {
	fmt.Println("🚀 Welcome to Task Tracker CLI! 🚀")

	store := filestore.NewTaskStore()

	cli := cli.InitCLI(store)
	cli.Run()
}
