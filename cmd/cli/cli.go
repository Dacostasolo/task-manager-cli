package cli

import (
	"flag"
	"os"

	"github.com/task-manager-cli/cmd/taskstore"
)

type CLI struct {
	store taskstore.Store
}

func (cli *CLI) dispatchCommand() {
	if len(os.Args) < 2 {
		cliUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		cli.addAction()
	case "list":
		cli.listAction()
	case "update":
		cli.updateAction()
	case "delete":
		cli.deleteAction()
	case "get":
		cli.getAction()
	case "persist":
		cli.persistAction()
	default:
		cliUsage()
	}
}

func InitCLI(store taskstore.Store) *CLI {
	flag.Usage = cliUsage

	return &CLI{store: store}
}

func (cli *CLI) Run() {
	cli.dispatchCommand()
}