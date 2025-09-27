package commands

import (
	"fmt"
	"lorech/advent-of-code/pkg/aoc"
	"lorech/advent-of-code/pkg/cmath"
	"lorech/advent-of-code/pkg/runners"

	"github.com/spf13/cobra"
)

func NewRandomCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Pick a random puzzle",
		RunE: func(cmd *cobra.Command, args []string) error {
			var randY, randD int

			o, _ := cmd.Flags().GetBool("open")
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

			if o {
				url, err := aoc.PuzzleUrl(randY, randD)
				if err != nil {
					return err
				}
				runners.OpenURL(url)
			} else {
				fmt.Printf("Day %d of %d", randD, randY)
			}

			return nil
		},
	}

	cmd.Flags().BoolP("open", "o", false, "open the puzzle instead of printing it")
	cmd.Flags().IntP("day", "d", 0, "lock in the day, picking between a random year's puzzle")
	cmd.Flags().IntP("year", "y", 0, "lock in the year, picking between it's puzzles")
	cmd.MarkFlagsMutuallyExclusive("day", "year")

	return cmd
}
