package aoc2025

import (
	"lorech/advent-of-code/pkg/cmath"
	"lorech/advent-of-code/pkg/convert"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type machine struct {
	lightsN int   // Total number of lights, used for targeting binary
	lights  int   // Target light state in binary
	buttons []int // Button wiring diagram in binary
	joltage []int // Target joltage
}

type lightState struct {
	depth  int // The depth of the current state, used for getting the total
	lights int // The current light state in binary
}

// Day 10: Factory
// https://adventofcode.com/2025/day/10
func dayTen(input string) (int, int) {
	return d10p1(input), d10p2(input)
}

// Completes the first half of the puzzle for day 10.
func d10p1(input string) int {
	factory := parseFactory(input)
	total := 0

	for _, m := range factory {
		s := lightState{depth: 0, lights: 0}

		stack, visited := []lightState{s}, []int{}
		for len(stack) != 0 {
			s := stack[0]
			stack = stack[1:]
			visited = append(visited, s.lights)

			if s.lights == m.lights {
				total += s.depth
				break
			}

			for _, b := range m.buttons {
				next := s.lights ^ b
				if !slices.Contains(visited, next) {
					new := lightState{depth: s.depth + 1, lights: next}
					if !slices.Contains(stack, new) {
						stack = append(stack, new)
					}
				}
			}
		}
	}

	return total
}

// Completes the second half of the puzzle for day 10.
// Reimplements u/Cute-Document3286's Zig solution as a learning excercise.
// https://github.com/gabrielmougard/AoC-2025/blob/main/10-factory/main.zig
func d10p2(input string) int {
	machines := parseFactory(input)
	totalCost := 0

	for _, m := range machines {
		totalCost += solvePartTwoMachine(m)
	}

	return totalCost
}

// Solves a single machine for the purposes of part two, where the machine
// input gets converted into a system of linear equations with some small
// brute-force through the possible value space, before arriving to a conclusion.
func solvePartTwoMachine(m machine) int {
	numButtons, numCounters := len(m.buttons), len(m.joltage)

	// Convert the machine to a matrix
	matrix := make([][]cmath.Rational, numCounters)
	for i := range matrix {
		matrix[i] = make([]cmath.Rational, numButtons+1)
		for j := range numButtons {
			matrix[i][j] = cmath.NewRationalInteger(0)
		}
		matrix[i][numButtons] = cmath.NewRationalInteger(m.joltage[i])
	}
	for i, b := range m.buttons {
		bb := convert.BinToIntIndex(b, len(m.joltage))
		for _, j := range bb {
			matrix[j][i] = cmath.NewRationalInteger(1)
		}
	}

	// Prepare for row reduction by initializing lookup values
	cr, pivotCols := 0, make([]int, len(matrix))
	for i := range pivotCols {
		pivotCols[i] = -1
	}

	// Pivot and row-reduce the matrix
	for ci := range numButtons {
		// Find the pivot row for this column
		pivotFound := false
		for ri := cr; ri < numCounters; ri++ {
			if !matrix[ri][ci].Equal(cmath.NewRationalInteger(0)) {
				if ri != cr {
					temp := matrix[cr]
					matrix[cr] = matrix[ri]
					matrix[ri] = temp
				}
				pivotFound = true
				break
			}
		}

		// Perform row reduction if a pivot was found
		if pivotFound {
			// Reduce the pivot row
			factor := matrix[cr][ci]
			for i := range numButtons + 1 {
				matrix[cr][i] = matrix[cr][i].Div(factor)
			}

			// Reduce the other rows
			for i := range numCounters {
				if i == cr {
					continue
				}

				factor = matrix[i][ci]
				for j := range numButtons + 1 {
					matrix[i][j] = matrix[i][j].Sub(matrix[cr][j].Mul(factor))
				}
			}

			pivotCols[cr] = ci
			cr++
		}
	}

	// Validate that the machine can even be solved by making sure
	// that any non-pivoted row does not expect a solution, which
	// would be impossible as each button press would need to be 0
	numPivots := cr
	for r := numPivots; r < numCounters; r++ {
		if !matrix[r][numButtons].Equal(cmath.NewRationalInteger(0)) {
			return 0
		}
	}

	// Determine which columns are pivots
	isPivot := make([]bool, len(m.buttons))
	for r := range cr {
		if pivotCols[r] >= 0 {
			isPivot[pivotCols[r]] = true
		}
	}

	// Collect free variables
	freeVars, numFree := []int{}, 0
	for c := range numButtons {
		if !isPivot[c] {
			freeVars = append(freeVars, c)
			numFree++
		}
	}

	// Identify the maximum number of presses for any given button
	maxCost := 0
	for _, j := range m.joltage {
		if j > maxCost {
			maxCost = j
		}
	}

	// Initialize moving values that are used for validation and backgracking
	bound, freeVals, minCost := maxCost+1, make([]int, numButtons), math.MaxInt

	// Brute-force every free variable value within the bounds of maximum cost
	solveFreeVars(
		&matrix,
		&pivotCols,
		&freeVars,
		&freeVals,
		&minCost,
		numButtons,
		numPivots,
		numFree,
		bound,
		0,
		0,
	)
	if minCost == math.MaxInt {
		return 0
	}
	return minCost
}

// Solves a single instance of the system of linear equations by inserting
// values into free variables recursively, until eventually a solution is
// found and compared against the running minimum total.
func solveFreeVars(
	matrix *[][]cmath.Rational,
	pivotCols *[]int,
	freeVars *[]int,
	freeVals *[]int,
	minCost *int,
	numButtons int,
	numPivots int,
	numFree int,
	bound int,
	depth int,
	freeCost int,
) {
	// No free cost remaining - minCost cannot decrease further
	if freeCost >= *minCost {
		return
	}

	// All free variables have values, so we have a solution
	if depth == numFree {
		// Initialize the solution with our free variables set to
		// their respective values, as determined by the recursion
		solution := make([]int, numButtons)
		for f := range numFree {
			solution[(*freeVars)[f]] = (*freeVals)[f]
		}

		cost, valid := freeCost, true

		// Back-substitute to solve pivot variables from bottom up
		ri := numPivots
		for ri > 0 {
			ri--
			col, val := (*pivotCols)[ri], (*matrix)[ri][numButtons]
			for c := col + 1; c < numButtons; c++ {
				val = val.Sub((*matrix)[ri][c].Mul(cmath.NewRationalInteger(solution[c])))
			}

			// Buttons cannot be pressed fractional number of times
			valInt, isInt := val.ToInt()
			if !isInt {
				valid = false
				break
			}

			// Buttons cannot be pressed negative number of times
			if valInt < 0 {
				valid = false
				break
			}

			solution[col] = valInt
			cost += valInt

			// Early exit, as this cannot be a better solution
			if cost >= *minCost {
				valid = false
				break
			}
		}

		// Found a valid solution which is smaller than our best one!
		if valid && cost < *minCost {
			*minCost = cost
		}

		return
	}

	// What is the maximum number we can insert into this free variable
	// to still feasibly reach a lower cost for the solution?
	thisBound := int(math.Min(
		float64(bound),
		math.Max(float64(*minCost-freeCost), 0),
	))

	for v := range thisBound {
		(*freeVals)[depth] = v
		solveFreeVars(
			matrix,
			pivotCols,
			freeVars,
			freeVals,
			minCost,
			numButtons,
			numPivots,
			numFree,
			bound,
			depth+1,
			freeCost+v,
		)

		// If the current best solution is smaller than the next one,
		// no point in even checking it
		if *minCost <= freeCost+v {
			break
		}
	}
}

// Parses input data into structured data.
func parseFactory(input string) []machine {
	rows := strings.Split(input, "\n")
	machines := make([]machine, len(rows))
	pre := regexp.MustCompile(`\[([.#]+)\] (\(.*\) ?)+ {([\d,]+)}`)
	bre := regexp.MustCompile(`[\d,]+`)

	for i, r := range rows {
		// Extract individual parts
		ps := pre.FindStringSubmatch(r)

		// Persist target light state
		ls := ps[1]
		for _, l := range ls {
			machines[i].lights <<= 0x1
			machines[i].lightsN++
			if l == '#' {
				machines[i].lights |= 0x1
			}
		}

		// Persist button wiring diagram
		bcsvs := bre.FindAllString(ps[2], -1)
		machines[i].buttons = make([]int, len(bcsvs))
		for j, bcsv := range bcsvs {
			bs, p := strings.Split(bcsv, ","), 0
			for _, b := range bs {
				bn, _ := strconv.Atoi(b)
				machines[i].buttons[j] <<= bn - p
				machines[i].buttons[j] |= 0x1
				p = bn
			}
			if p < machines[i].lightsN-1 {
				machines[i].buttons[j] <<= machines[i].lightsN - 1 - p
			}
		}

		// Persist joltage targets
		jcsv := ps[3]
		js := strings.Split(jcsv, ",")
		machines[i].joltage = make([]int, len(js))
		for k, j := range js {
			jn, _ := strconv.Atoi(j)
			machines[i].joltage[k] = jn
		}
	}

	return machines
}
