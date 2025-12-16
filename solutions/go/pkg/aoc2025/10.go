package aoc2025

import (
	"fmt"
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

// 18864 < answer < 21055
// Not: 20653

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

		// Pivot and row-reduce the matrix
		cr, pivots := 0, map[int]int{}
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

				pivots[cr] = ci
				cr++
			}
		}

		// Identify the free variables in each equation
		freeVars, solvedVars := map[int][]int{}, map[int]int{}
		for pr, pc := range pivots {
			for i := pc + 1; i < len(m.buttons); i++ {
				if !matrix[pr][i].Equal(cmath.NewRationalInteger(0)) {
					freeVars[i] = append(freeVars[i], pr)
				}
			}
		}

		// Brute-force the free variables to find the solution
		minCost := math.MaxInt
		minCost, _ = solveFreeVars(matrix, solvedVars, freeVars, &minCost)
		totalCost += minCost
	}

	return totalCost
}

// Solves the `matrix` by inserting `freeVars` into the equation with the values in `solvedVars`.
func solveFreeVars(matrix [][]cmath.Rational, solvedVars map[int]int, freeVars map[int][]int, minCost *int) (int, bool) {
	valid, cost := true, 0
	// If there are no free variables, the equations are solved
	if len(freeVars) == 0 {
		for _, r := range matrix {
			c := rowCost(r, solvedVars)
			// Buttons cannot be pressed negative times, so the solution is invalid
			if c < 0 {
				valid = false
			}
			cost += c
		}
		// Add in the current values of the free variables
		// Since those are controlled, we know they bottom out at 0 - no need to check negatives
		for _, v := range solvedVars {
			cost += v
		}
		return cost, valid
	}

	// Get the button this instance of the function will increment
	keys := make([]int, 0, len(freeVars))
	for k := range freeVars {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	button := keys[0]

	// Backup the affected rows before removing the button from the equation
	// to set up a compatible base case for the recursion
	affected := freeVars[button]

	// Since joltage cannot be decreased, the worst case scenario number of button
	// presses for our free variables is equal to the total joltage of every total
	// that is affected by this button press
	// TODO: Can anything more efficient be used?
	maxPresses := 0
	for i, r := range matrix {
		if slices.Contains(affected, i) {
			num := r[len(matrix[0])-1]
			val := num.Numerator
			if val > maxPresses {
				maxPresses = val
			}
		}
	}

	// Find the minimum cost while iterating through options for this free variable
	delete(freeVars, button)
	pCost := math.MaxInt
	for n := range maxPresses {
		// Insert the new variable into the equation and solve the equation
		old, used := solvedVars[button]
		solvedVars[button] = n
		cost, valid := solveFreeVars(matrix, solvedVars, freeVars, minCost)

		// Minimum cost can only decrease if the result is valid
		if valid {
			if cost < *minCost {
				*minCost = cost
				break
			}
		}

		// Restore the previous value to prevent cross-contamination
		if used {
			solvedVars[button] = old
		} else {
			delete(solvedVars, button)
		}

		// Prune values that will increase the final cost if we have found
		// some value for the variable that actually produces a valid cost
		if pCost < math.MaxInt && cost-pCost > 0 {
			break
		}
		pCost = cost
	}
	freeVars[button] = affected

	return *minCost, valid
}

// Computes the right-hand side result (total cost) of a row from the reduced
// matrix given a solved value for each free variable in the matrix.
func rowCost(row []cmath.Rational, solvedVars map[int]int) int {
	result := row[len(row)-1]

	for col, val := range solvedVars {
		factor := row[col]
		result = result.Sub(factor.Mul(cmath.NewRationalInteger(val)))
	}

	return result.Numerator
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
