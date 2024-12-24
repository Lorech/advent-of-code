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
func dayTwentyFour(input string) (int, int) {
	return d24p1(input), 0
}

// Completes the first half of the puzzle for day 24.
func d24p1(input string) int {
	system := parseGates(input)

	// Find all the z-gates of this particicular system.
	zs := make([]string, 0)
	for gate, _ := range system {
		if gate[0] == 'z' {
			zs = append(zs, gate)
		}
	}

	// Descending sort: maps have undefined order, so we must sort it anyway, and
	// the puzzle puzzle puts the last bit into z00, meaning we need descending sort.
	slices.SortFunc(zs, func(a, b string) int {
		return -strings.Compare(a, b)
	})

	nums := make([]int, len(zs))
	for i, z := range zs {
		num := system[z].val(system)
		nums[i] = num
	}
	result, _ := convert.Stoi(nums, 2)
	return result
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
