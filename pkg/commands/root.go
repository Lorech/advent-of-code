package commands

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aoc",
		Short: "Advent of Code CLI",
		Long:  "Lightweight utility for quickly setting up for Advent of Code",
	}

	cmd.AddCommand(NewVersionCommand())
	cmd.AddCommand(NewOpenCommand())

	return cmd
}
