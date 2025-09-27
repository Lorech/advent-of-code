package commands

import (
	"fmt"
	"lorech/advent-of-code/pkg/aoc"
	"lorech/advent-of-code/pkg/cmath"

	"github.com/spf13/cobra"
)

func NewRandomCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Pick a random puzzle",
		Run: func(cmd *cobra.Command, args []string) {
			var randY, randD int

			y, _ := cmd.Flags().GetInt("year")
			d, _ := cmd.Flags().GetInt("day")

			if y == 0 {
				randY = cmath.RandomInRange(aoc.MinYear, aoc.MaxYear())
			} else {
				randY = y
			}

			if d == 0 {
				randD = cmath.RandomInRange(aoc.MinDay, aoc.MaxDay)
			} else {
				randD = d
			}

			fmt.Printf("Day %d of %d", randD, randY)
		},
	}

	cmd.Flags().IntP("day", "d", 0, "lock in the day, picking between a random year's puzzle")
	cmd.Flags().IntP("year", "y", 0, "lock in the year, picking between it's puzzles")
	cmd.MarkFlagsMutuallyExclusive("day", "year")

	return cmd
}
