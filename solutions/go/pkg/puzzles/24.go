package puzzles

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/convert"
	"slices"
	"strconv"
	"strings"
)

type operation int

const (
	NONE operation = iota
	AND
	XOR
	OR
)

type gate struct {
	name  string    // The identifying name of the gate.
	value int       // The current value of the gate; 0 or 1 if known, -1 if not.
	a     string    // The first operand in the value's calculation.
	b     string    // The second operand in the value's calculation.
	op    operation // The operation to apply to both operands to get the value.
}

func newKnownGate(name string, value int) gate {
	var g gate
	g.name = name
	g.value = value
	return g
}

func newCalculatedGate(name string, a string, b string, op operation) gate {
	var g gate
	g.name = name
	g.value = -1
	g.a = a
	g.b = b
	g.op = op
	return g
}

func (g gate) val(lookup map[string]gate) int {
	if g.value != -1 {
		return g.value
	}

	a, b := convert.Itob(lookup[g.a].val(lookup)), convert.Itob(lookup[g.b].val(lookup))
	v := 0
	switch g.op {
	case AND:
		v = convert.Btoi(a && b)
	case OR:
		v = convert.Btoi(a || b)
	case XOR:
		v = convert.Btoi(a != b)
	default:
		panic(fmt.Sprintf("Gate %s has an invalid operation!", g.name))
	}

	g.value = v
	return v
}

// Day 24: Crossed Wires
// https://adventofcode.com/2024/day/24
func dayTwentyFour(input string) (int, string) {
	return d24p1(input), d24p2(input)
}

// Completes the first half of the puzzle for day 24.
func d24p1(input string) int {
	system := parseGates(input)
	numbers := extractNumbers(system, 'z')
	result, _ := convert.Stoi(numbers, 2)
	return result
}

// Completes the second half of the puzzle for day 24.
func d24p2(input string, options ...int) string {
	system := parseGates(input)
	zs := extractNumbers(system, 'z')

	wrongs := make([]string, 0)
	for output, gate := range system {
		// This is an output of the bit adder.
		if output[0] == 'z' {
			// The sum value can only be created by an XOR operation, unless it's the
			// last bit, in which case it's just a carry bit.
			if gate.op != XOR {
				if output != fmt.Sprintf("z%02d", len(zs)-1) {
					wrongs = append(wrongs, output)
				}

				// We already checked if this is an invalid output or not, so move on.
				continue
			}

			// This is a valid output, but it may have invalid inputs.
			operands := [2]string{gate.a, gate.b}
			for _, operand := range operands {
				parent := system[operand]

				// This is an input value, so it can't be incorrect.
				if parent.op == NONE {
					continue
				}

				// This is the result of the sum block from a full adder, so it must be
				// created with x or y, with the other number being the carry. Logic is
				// non-commutative, so A will always store the x/y if it is present.
				if parent.op == XOR && !(strings.HasPrefix(parent.a, "x") || strings.HasPrefix(parent.a, "y")) {
					wrongs = append(wrongs, operand)
				}

				// This is the result of the sum block from a half adder, so it must be
				// the first bit of the output.
				if parent.op == AND && !strings.HasSuffix(parent.a, "00") {
					wrongs = append(wrongs, operand)
				}
			}
		}

		// This is a carry bit.
		if gate.op == OR {
			operands := [2]string{gate.a, gate.b}
			for _, operand := range operands {
				// A carry bit can only be created by an AND operation.
				if system[operand].op != AND {
					wrongs = append(wrongs, operand)
				}
			}
		}
	}

	slices.Sort(wrongs)
	return strings.Join(wrongs, ",")
}

// Finds all the numbers starting within the start rune in the lookup table,
// constructing a slice of their values in binary form.
func extractNumbers(system map[string]gate, start rune) []int {
	// Find all the gates of this particicular system.
	gates := make([]string, 0)
	for gate, _ := range system {
		if rune(gate[0]) == start {
			gates = append(gates, gate)
		}
	}

	// Descending sort: maps have undefined order, so we must sort it anyway, and
	// the puzzle puzzle puts the last bit into 00, meaning we need descending sort.
	slices.SortFunc(gates, func(a, b string) int {
		return -strings.Compare(a, b)
	})

	nums := make([]int, len(gates))
	for i, gate := range gates {
		num := system[gate].val(system)
		nums[i] = num
	}

	return nums
}

// Parses the input data into structured data for gate value resolval.
func parseGates(input string) map[string]gate {
	rows := strings.Split(input, "\n")
	gates := make(map[string]gate)
	for _, row := range rows {
		// Delimiter between data sets.
		if len(row) == 0 {
			continue
		}

		data := strings.Split(row, " ")
		switch len(data) {
		case 2: // Starting values.
			name := data[0][:len(data[0])-1]
			value, _ := strconv.Atoi(data[1])
			gates[name] = newKnownGate(name, value)
		case 5: // Gate logic.
			op := NONE
			if data[1] == "AND" {
				op = AND
			} else if data[1] == "XOR" {
				op = XOR
			} else if data[1] == "OR" {
				op = OR
			}
			gates[data[4]] = newCalculatedGate(data[4], data[0], data[2], op)
		}
	}
	return gates
}
