package aoc2025

import (
	"regexp"
	"strconv"
	"strings"
)

type calculation struct {
	nums []string
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
		s, _ := strconv.Atoi(strings.Trim(c.nums[0], " "))
		for _, num := range c.nums[1:] {
			n, _ := strconv.Atoi(strings.Trim(num, " "))
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

	// Starting indices of each unique calculation
	ops := regexp.MustCompile(`[+*]`)
	starts := ops.FindAllStringIndex(rows[len(rows)-1], -1)

	// Allocate memory for all calculations based on the number
	// number of operators (calculations) and rows (operands)
	calcs := make([]calculation, len(starts))
	for i := range calcs {
		calcs[i].nums = make([]string, len(rows)-1)
	}

	// Extract all the operators for each calculation
	for i, s := range starts {
		calcs[i].op = rows[len(rows)-1][s[0]:s[1]]
	}

	// Extract all the operands for each calculation
	for i, r := range rows[:len(rows)-1] {
		for j := range starts {
			s, e := starts[j][0], len(r)
			if j < len(starts)-1 {
				e = starts[j+1][0] - 1
			}
			calcs[j].nums[i] = r[s:e]
		}
	}

	return calcs
}
