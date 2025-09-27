package commands

import (
	"lorech/advent-of-code/pkg/aoc"
	"lorech/advent-of-code/pkg/runners"

	"github.com/spf13/cobra"
)

func NewOpenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the puzzle of the day",
		RunE: func(cmd *cobra.Command, args []string) error {
			year, _ := cmd.Flags().GetInt("year")
			day, _ := cmd.Flags().GetInt("day")
			input, _ := cmd.Flags().GetBool("input")

			var url string
			var err error
			if input {
				url, err = aoc.InputUrl(year, day)
			} else {
				url, err = aoc.PuzzleUrl(year, day)
			}

			if err != nil {
				return err
			}

			runners.OpenURL(url)
			return nil
		},
	}

	cmd.Flags().IntP("day", "d", aoc.ClosestDay(), "day to open")
	cmd.Flags().IntP("year", "y", aoc.MaxYear(), "year to open")
	cmd.Flags().BoolP("input", "i", false, "open the input")

	return cmd
}
