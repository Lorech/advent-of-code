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
	eActuallySafe := 4

	if rSafe, rActuallySafe := dayTwo(input); eSafe != rSafe || eActuallySafe != rActuallySafe {
		t.Errorf("DayTwo() = %v, %v, want %v, %v", rSafe, rActuallySafe, eSafe, eActuallySafe)
	}
}

// Tests a few edge cases that helped narrow down the solution for part 2.
func TestDayOneEdgeCases(t *testing.T) {
	input := `9 1 2 3 4
1 2 3 4 9
2 1 3 5 8
2 6 1`

	eSafe := 0
	eActuallySafe := 4

	if rSafe, rActuallySafe := dayTwo(input); eSafe != rSafe || eActuallySafe != rActuallySafe {
		t.Errorf("DayTwo() = %v, %v, want %v, %v", rSafe, rActuallySafe, eSafe, eActuallySafe)
	}
}
