package commands

import (
	"fmt"
	"lorech/advent-of-code/pkg/cmath"
	"lorech/advent-of-code/pkg/runners"
	"time"

	"github.com/spf13/cobra"
)

const actualMinDay = 1
const actualMaxDay = 31
const minPuzzleDay = 1
const maxPuzzleDay = 25
const minPuzzleYear = 2015

func NewOpenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the puzzle of the day",
		Args: func(cmd *cobra.Command, args []string) error {
			_, maxPuzzleYear := defaultPuzzle()
			year, _ := cmd.Flags().GetInt("year")
			if year < minPuzzleYear || year > maxPuzzleYear {
				return fmt.Errorf("invalid year specified: %d", year)
			}

			day, _ := cmd.Flags().GetInt("day")
			if day < minPuzzleDay || day > maxPuzzleDay {
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

	defaultDay, defaultYear := defaultPuzzle()
	cmd.Flags().IntP("day", "d", defaultDay, "day to open")
	cmd.Flags().IntP("year", "y", defaultYear, "year to open")

	return cmd
}

// Find the default puzzle.
func defaultPuzzle() (int, int) {
	now := time.Now()

	day := cmath.ClosestInRange(now.Day(), 1, 31, minPuzzleDay, maxPuzzleDay)

	var year int
	if now.Month() == 12 {
		year = now.Year()
	} else {
		year = now.Year() - 1
	}

	return day, year
}
