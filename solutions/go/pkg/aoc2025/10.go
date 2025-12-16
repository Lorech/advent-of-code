package aoc2025

import (
	"fmt"
	"lorech/advent-of-code/pkg/cmath"
	"lorech/advent-of-code/pkg/convert"
	"maps"
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

// 18864 < answer < 21055
// Not: 19900, 20653

// Completes the second half of the puzzle for day 10.
func d10p2(input string) int {
	machines := parseFactory(input)
	totalCost := 0

	for i, m := range machines {
		fmt.Printf("Machine %d:\n", i+1)
		// Convert the machine to a matrix
		matrix := make([][]cmath.Rational, len(m.joltage))
		for i := range matrix {
			matrix[i] = make([]cmath.Rational, len(m.buttons)+1)
			for j := range len(m.buttons) {
				matrix[i][j] = cmath.NewRationalInteger(0)
			}
			matrix[i][len(m.buttons)] = cmath.NewRationalInteger(m.joltage[i])
		}
		for i, b := range m.buttons {
			bb := convert.BinToIntIndex(b, len(m.joltage))
			for _, j := range bb {
				matrix[j][i] = cmath.NewRationalInteger(1)
			}
		}

		cr, pivotCols := 0, make([]int, len(matrix))
		for i := range pivotCols {
			pivotCols[i] = -1
		}

		// Pivot and row-reduce the matrix
		for ci := range len(m.buttons) {
			// Find the pivot row for this column
			pivotFound := false
			for ri := cr; ri < len(matrix); ri++ {
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
				for i := range len(m.buttons) + 1 {
					matrix[cr][i] = matrix[cr][i].Div(factor)
				}

				// Reduce the other rows
				for i := range len(matrix) {
					if i == cr {
						continue
					}

					factor = matrix[i][ci]
					for j := range len(m.buttons) + 1 {
						matrix[i][j] = matrix[i][j].Sub(matrix[cr][j].Mul(factor))
					}
				}

				pivotCols[cr] = ci
				cr++
			}
		}

		hasSolution := true
		for r := cr; r < len(matrix); r++ {
			if !matrix[r][len(matrix[0])-1].Equal(cmath.NewRationalInteger(0)) {
				hasSolution = false
			}
		}
		if !hasSolution {
			continue
		}

		// Identify the free variables in each equation
		pivotColsByRow := make([]int, len(matrix))
		for i := range pivotColsByRow {
			pivotColsByRow[i] = -1
		}
		for r := 0; r < cr; r++ { // cr = number of pivot rows after elimination
			pivotColsByRow[r] = pivotCols[r] // pivotCols[r] was set during elimination
		}

		// Determine which columns are pivots
		isPivot := make([]bool, len(m.buttons))
		for _, col := range pivotColsByRow {
			if col >= 0 {
				isPivot[col] = true
			}
		}

		// Collect free variables
		freeVars := []int{}
		for c := 0; c < len(m.buttons); c++ {
			if !isPivot[c] {
				freeVars = append(freeVars, c)
			}
		}

		// Initialize solved variables map
		solvedVars := map[int]int{}

		// Identify the maximum number of presses for any given button
		maxCost := 0
		for _, j := range m.joltage {
			if j > maxCost {
				maxCost = j
			}
		}

		// Brute-force the free variables to find the solution
		minCost := math.MaxInt
		minCost, _ = solveFreeVars(matrix, pivotColsByRow, freeVars, solvedVars, 0, &minCost, maxCost)
		totalCost += minCost
	}

	return totalCost
}

// Solves the `matrix` by inserting `freeVars` into the equation with the values in `solvedVars`.
func solveFreeVars(
	matrix [][]cmath.Rational,
	pivotCols []int,
	freeVars []int,
	solvedVars map[int]int,
	depth int,
	minCost *int,
	maxCost int,
) (int, bool) {
	valid, cost := true, 0
	// If there are no free variables, the equations are solved
	if depth == len(freeVars) {
		solution := make(map[int]int)
		maps.Copy(solution, solvedVars)
		for _, v := range solvedVars {
			cost += v
		}

		for pr := len(pivotCols) - 1; pr >= 0; pr-- {
			col := pivotCols[pr]
			if col < 0 {
				continue
			}
			rhs := matrix[pr][len(matrix[0])-1]

			for c := col + 1; c < len(matrix[0])-1; c++ {
				if val, ok := solution[c]; ok {
					rhs = rhs.Sub(matrix[pr][c].Mul(cmath.NewRationalInteger(val)))
				}
			}

			if !rhs.Integer() || rhs.Numerator < 0 {
				return math.MaxInt, false
			}

			v := rhs.Numerator
			solution[col] = v
			cost += v

			if cost >= *minCost {
				return math.MaxInt, false
			}
		}

		return cost, valid
	}

	// Find the free variable this function instance will iterate over
	button := freeVars[depth]
	remaining := *minCost
	for _, v := range solvedVars {
		remaining -= v
	}
	if remaining < 0 {
		return math.MaxInt, false
	}
	maxPresses := maxCost
	if remaining < maxPresses {
		maxPresses = remaining
	}

	// Iterate through the first free variable, finding minimum cost
	for n := 0; n <= maxPresses; n++ {
		// Insert the new variable into the equation and solve the equation
		old, used := solvedVars[button]
		solvedVars[button] = n
		cost, valid := solveFreeVars(
			matrix,
			pivotCols,
			freeVars,
			solvedVars,
			depth+1,
			minCost,
			maxCost,
		)

		// Minimum cost can only decrease if the result is valid
		if valid && cost < *minCost {
			*minCost = cost
			break
		}

		// Restore the previous value to prevent cross-contamination
		if used {
			solvedVars[button] = old
		} else {
			delete(solvedVars, button)
		}
	}

	return *minCost, valid
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
