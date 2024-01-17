package main

import (
	"flag"
	"fmt"

	"github.com/joaogabsoaresf/reps/commands"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		commands.RunHelpCommand()
		// flag.Usage()
		return
	}

	command := args[0]

	var commandArgs string

	if len(args) > 1 {
		commandArgs = args[1]
	}

	switch command {
	case "init":
		commands.RunInitCommand(commandArgs)
	default:
		fmt.Println("Unknown command:", command)
	}
}
