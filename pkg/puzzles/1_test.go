package puzzles

import "testing"

// Tests the full solution for day 1 based on the provided example.
func TestDayOne(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	eDistance := 11
	eSimilarity := 31

	if rDistance, rSimilarity := dayOne(input); eDistance != rDistance || eSimilarity != rSimilarity {
		t.Errorf("DayOne() = %v, %v, want %v, %v", rDistance, rSimilarity, eDistance, eSimilarity)
	}
}
