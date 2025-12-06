package aoc2025

import (
	"regexp"
	"strconv"
	"strings"
)

type calculation struct {
	nums []int
	op   string
}

// Day 6: Trash Compactor
// https://adventofcode.com/2025/day/6
func daySix(input string) (int, int) {
	return d6p1(input), 0
}

// Completes the first half of the puzzle for day 6.
func d6p1(input string) int {
	calcs := parseHomework(input)
	sum := 0

	for _, c := range calcs {
		s := c.nums[0]
		for _, n := range c.nums[1:] {
			if c.op == "+" {
				s += n
			} else if c.op == "*" {
				s *= n
			}
		}
		sum += s
	}

	return sum
}

func parseHomework(input string) []calculation {
	rows := strings.Split(input, "\n")
	digits := regexp.MustCompile(`\d+`)

	// Allocate memory for the number of calculations (horizontal)
	// and the number of digits in the calculations (vertical)
	n := len(digits.FindAllString(rows[0], -1))
	calcs := make([]calculation, n)
	for i := range calcs {
		calcs[i].nums = make([]int, len(rows)-1)
	}

	// Build the calculations by removing all repeating spaces,
	// and inserting them into the calculations by rotating the 2D array.
	spaces := regexp.MustCompile(`[ ]+`)
	for i, r := range rows {
		r = spaces.ReplaceAllString(strings.Trim(r, " "), " ")
		p := strings.Split(r, " ")
		for j, d := range p {
			if i == len(rows)-1 {
				calcs[j].op = d
			} else {
				num, _ := strconv.Atoi(d)
				calcs[j].nums[i] = num
			}
		}
	}

	return calcs
}
