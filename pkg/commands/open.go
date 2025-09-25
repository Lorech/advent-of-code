package commands

import (
	"fmt"
	"lorech/advent-of-code/pkg/aoc"
	"lorech/advent-of-code/pkg/runners"

	"github.com/spf13/cobra"
)

func NewOpenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the puzzle of the day",
		Args: func(cmd *cobra.Command, args []string) error {
			year, _ := cmd.Flags().GetInt("year")
			if year < aoc.MinYear || year > aoc.MaxYear() {
				return fmt.Errorf("invalid year specified: %d", year)
			}

			day, _ := cmd.Flags().GetInt("day")
			if day < aoc.MinDay || day > aoc.MaxDay {
				return fmt.Errorf("invalid day specified: %d", day)
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			year, _ := cmd.Flags().GetInt("year")
			day, _ := cmd.Flags().GetInt("day")
			runners.OpenURL(fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day))
		},
	}

	cmd.Flags().IntP("day", "d", aoc.ClosestDay(), "day to open")
	cmd.Flags().IntP("year", "y", aoc.MaxYear(), "year to open")

	return cmd
}
