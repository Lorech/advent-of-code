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
	want := 11

	if got := DayOne(input); got != want {
		t.Errorf("DayOne() = %v, want %v", got, want)
	}
}
