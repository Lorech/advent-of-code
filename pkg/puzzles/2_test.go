package puzzles

import (
	"lorech/advent-of-code-2024/pkg/file"
	"testing"
)

// Tests the full solution for day 2 based on the provided example.
func TestDayTwo(t *testing.T) {
	input, err := file.ReadTestFile(2)

	if err != nil {
		t.Errorf("Could not read test file: %v", err)
		return
	}

	eSafe := 2
	eActuallySafe := 4

	if rSafe, rActuallySafe := dayTwo(input); eSafe != rSafe || eActuallySafe != rActuallySafe {
		t.Errorf("DayTwo() = %v, %v, want %v, %v", rSafe, rActuallySafe, eSafe, eActuallySafe)
	}
}
