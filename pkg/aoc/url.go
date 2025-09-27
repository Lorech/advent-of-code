package aoc

import "fmt"

// Attempts to generate the URL to a puzzle, returning it if successful.
func PuzzleUrl(year, day int) (string, error) {
	if year < MinYear || year > MaxYear() {
		return "", fmt.Errorf("invalid year specified: %d", year)
	}

	if day < MinDay || day > MaxDay {
		return "", fmt.Errorf("invalid day specified: %d", day)
	}

	return fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day), nil
}

// Attempts to generate the URL to a puzzle input, returning it if successful.
func InputUrl(year, day int) (string, error) {
	url, err := PuzzleUrl(year, day)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/input", url), nil
}
