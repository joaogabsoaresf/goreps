package commands

import (
	"fmt"
	"os"
	"strings"
)

func initUsage() string {
	return `Usage:

	rep init [path]

The path argument can be blank to create .rep
in the current directory, or can be passed a
argument with the path in the format: ~/dir/dir2/dir3.
`
}

func RunInitCommand(args ...string) {
	if len(args) > 1 {
		fmt.Println(initUsage())
	} else if len(args) == 0 || args[0] == "." || args[0] == "" {
		createDotRep(".")
	} else {
		if !strings.Contains(args[0], "~/") {
			fmt.Print("Args must be a path or .\n")
			fmt.Println(initUsage())
		} else {
			createDotRep(args[0])
		}
	}
}

func createDotRep(path string) error {
	var fPath string
	if path != "." {
		fPath = fmt.Sprintf("/%s/.rep", path)
	} else {
		fPath = ".rep"
	}
	err := os.Mkdir(fPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Error to initialize rep: %v", err)
		return err
	}
	fmt.Printf("Rep created successfully in %s!", path)
	return nil
}
