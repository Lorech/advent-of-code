package puzzles

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type clawMachine struct {
	A clawButton // The configuration of button A.
	B clawButton // The configuration of button B.
	P clawPrize  // The parameters of the prize.
}

type clawButton struct {
	X    int // The number of units moved to the right.
	Y    int // The number of units moved upward.
	Cost int // The cost of pressing the button.
}

type clawPrize struct {
	X int // The position of the prize on the x-axis.
	Y int // The position of the prize on the y-axis.
}

// Day 13: Claw Contraption
// https://adventofcode.com/2024/day/13
func dayThirteen(input string) (int, int) {
	return d13p1(input), 0
}

// Completes the first half of the puzzle for day 13.
func d13p1(input string) int {
	machines := parseClawMachines(input)
	price := 0

	for _, machine := range machines {
		cost, valid := calculateCost(machine.A, machine.B, machine.P)
		if valid {
			price += cost
		}
	}

	return price
}

// Calculates the minimum cost to position the claw over the prize. Returns
// this cost, and boolean determining if this cost should be considered valid,
// as math.MaxInt gets returned otherwise.
func calculateCost(a clawButton, b clawButton, r clawPrize) (int, bool) {
	cost := math.MaxInt

	// Keep the number of iterations as low as possible - 100, as per the limits
	// of the puzzle, or less, if it would overflow the result, creating
	// unnecessary calculations.
	aMaxX := (r.X + a.X - 1) / a.X
	bMaxX := (r.X + b.X - 1) / b.X
	aMaxY := (r.Y + a.Y - 1) / a.Y
	bMaxY := (r.Y + b.Y - 1) / b.Y

	for x := range min(100, aMaxX, aMaxY) {
		for y := range min(100, bMaxX, bMaxY) {
			vX := x*a.X + y*b.X
			vY := x*a.Y + y*b.Y

			// Early exit, since we overflowed the result and will never match it.
			if vX > r.X && vY > r.Y {
				break
			}

			// Valid move!
			if vX == r.X && vY == r.Y && a.Cost*x+b.Cost*y < cost {
				cost = a.Cost*x + b.Cost*y
			}
		}
	}

	return cost, cost != math.MaxInt
}

// Parses the input string into structured data.
func parseClawMachines(input string) []clawMachine {
	rows := strings.Split(input, "\n")
	machines := make([]clawMachine, 0)
	machine := clawMachine{}
	reDigits := regexp.MustCompile(`\d+`)

	for i, row := range rows {
		switch i % 4 {
		case 0:
			values := reDigits.FindAllString(row, -1)
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(values[1])
			machine.A = clawButton{x, y, 3}
		case 1:
			values := reDigits.FindAllString(row, -1)
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(values[1])
			machine.B = clawButton{x, y, 1}
		case 2:
			values := reDigits.FindAllString(row, -1)
			x, _ := strconv.Atoi(values[0])
			y, _ := strconv.Atoi(values[1])
			machine.P = clawPrize{x, y}
		case 3:
			machines = append(machines, machine)
		}
	}

	// Append the last machine to the row.
	// WARNING: If the input ends with a newline, this will add the last machine
	// twice, but for now it doesn't so I don't care to fix it!
	machines = append(machines, machine)

	return machines
}
