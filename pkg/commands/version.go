package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the running CLI's version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("aoc v0.1.0")
		},
	}
}
