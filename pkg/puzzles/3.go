package puzzles

import (
	"regexp"
	"strconv"
)

// Day 3: Mull It Over
// https://adventofcode.com/2024/day/3
func dayThree(input string) (int, int) {
	multiplication := 0

	re := regexp.MustCompile(`(mul\(\d+?,\d+?\))`)
	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		reDigits := regexp.MustCompile(`\d+`)
		digits := reDigits.FindAllString(match, -1)
		first, _ := strconv.Atoi(digits[0])
		second, _ := strconv.Atoi(digits[1])
		multiplication += first * second
	}

	return multiplication, 0
}
