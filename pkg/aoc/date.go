package aoc

import (
	"lorech/advent-of-code/pkg/cmath"
	"time"
)

const MinDay = 1
const MaxDay = 25
const MinYear = 2015

// Gets the maximum year that a puzzle can be picked from.
//
// Picks the current year during December; last year otherwise.
func MaxYear() int {
	now := time.Now()

	var year int
	if now.Month() == 12 {
		year = now.Year()
	} else {
		year = now.Year() - 1
	}

	return year
}

// Finds the closest puzzle day to current day in a circular fashion.
func ClosestDay() int {
	now := time.Now()
	return cmath.ClosestInRange(now.Day(), 1, 31, MinDay, MaxDay)
}
