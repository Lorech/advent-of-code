package main

import (
	"fmt"
	"lorech/advent-of-code/pkg/commands"
	"os"
)

func main() {
	cmd := commands.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
