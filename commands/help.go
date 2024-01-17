package commands

import "fmt"

func RunHelpCommand() {
	fmt.Println(helpString())
}

func helpString() string {
	return `Hello, this is the CLI of GO Reps. A CLI to manage your repositories.
Usage:

	rep <command> [arguments]

The available commands are:

	- help:	To see the list and explanation of all commands.
	- init: To initialize a new rep in your project.
	- add: 	To add files to the queue of changes.
	- commit: To save your changes in your local rep.
	- log: 	To list a resume of your last five commits,
	from the newest to the oldest.
`
}
