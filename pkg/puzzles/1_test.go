package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the full solution for day 1 based on the provided example.
func TestDayOne(t *testing.T) {
	input, err := file.ReadTestFile(1)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	eDistance := 11
	eSimilarity := 31

	if rDistance, rSimilarity := dayOne(input); eDistance != rDistance || eSimilarity != rSimilarity {
		t.Errorf("DayOne() = %v, %v, want %v, %v", rDistance, rSimilarity, eDistance, eSimilarity)
	}
}
