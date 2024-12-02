package puzzles

import "testing"

// Tests the full solution for day 2 based on the provided example.
func TestDayTwo(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	eSafe := 2

	if rSafe, _ := DayTwo(input); eSafe != rSafe {
		t.Errorf("DayTwo() = %v, %v, want %v, %v", rSafe, 0, eSafe, 0)
	}
}
