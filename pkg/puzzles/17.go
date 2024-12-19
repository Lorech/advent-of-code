package puzzles

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/cslices"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Day 17: Chronospacial Computer
// https://adventofcode.com/2024/day/17
func daySeventeen(input string) (interface{}, interface{}) {
	return d17p1(input), d17p2(input)
}

// Completes the first half of the puzzle for day 17.
func d17p1(input string) string {
	a, b, c, program := parseProgram(input)
	return runProgram(program, a, b, c)
}

// Completes the second half of the puzzle for day 17.
func d17p2(input string) int {
	_, _, _, program := parseProgram(input)
	solutions := make([]int, 0)
	reverseEngineer([]int{}, program, &solutions)
	slices.Sort(solutions)
	return solutions[0]
}

// Recursively reverse engineers the programt to find all solutions that output
// the same code as the program itself.
func reverseEngineer(digits []int, program []int, solutions *[]int) {
	// End of the recursion - time to check if the program is fully the same.
	if len(digits) == len(program) {
		a, _ := cslices.Stoi(digits, 8)
		r := runProgram(program, a, 0, 0)
		if match(program, r) {
			*solutions = append(*solutions, a)
		}
		return
	}

	// BFS through the possible values of the next digit.
	e := program[len(program)-len(digits)-1:]
	for i := range 8 {
		d := append(append([]int(nil), digits...), i)
		a, _ := cslices.Stoi(d, 8)
		r := runProgram(program, a, 0, 0)
		if match(e, r) {
			reverseEngineer(d, program, solutions)
		}
	}
}

// Compares an integer slice (source code) to a string (program output) to
// determine if the program generated itself.
func match(program []int, result string) bool {
	r := strings.Split(result, ",")

	if len(r) != len(program) {
		return false
	}

	for i := range r {
		if r[i] != strconv.Itoa(program[i]) {
			return false
		}
	}

	return true
}

// Runs the provided program, returning the string the program would output.
func runProgram(program []int, a int, b int, c int) string {
	var o strings.Builder
	ip := 0

	for ip < len(program)-1 {
		opcode, literal := program[ip], program[ip+1]

		switch opcode {
		case 0: // ADV
			a = a / int(math.Pow(2, float64(comboOperand(literal, a, b, c))))
		case 1: // BXL
			b = b ^ literal
		case 2: // BST
			b = comboOperand(literal, a, b, c) % 8
		case 3: // JNZ
			if a != 0 {
				ip = literal
				continue // The pointer does not move if we jump.
			}
		case 4: // BXC
			b = b ^ c
		case 5: // OUT
			v := comboOperand(literal, a, b, c) % 8
			if o.Len() != 0 {
				o.WriteString(fmt.Sprintf(",%d", v))
			} else {
				o.WriteString(fmt.Sprintf("%d", v))
			}
		case 6: // BDV
			b = a / int(math.Pow(2, float64(comboOperand(literal, a, b, c))))
		case 7: // CDV
			c = a / int(math.Pow(2, float64(comboOperand(literal, a, b, c))))
		}

		ip += 2
	}

	return o.String()
}

// Resolves the value of a combo operand based on the input literal and the
// current values of the computer's registers.
func comboOperand(literal int, a int, b int, c int) int {
	switch literal {
	case 0, 1, 2, 3:
		return literal
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	case 7:
		panic("You said 7 wouldn't appear!")
	default:
		fmt.Println(literal)
		panic("You didn't even mention this value!")
	}
}

// Parses the input data into values of the A, B, and C registers, and a slice
// containing the actual program code.
func parseProgram(input string) (int, int, int, []int) {
	var a, b, c int
	var code []int

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		// Delimiter between the initial registers and the program code.
		if len(row) == 0 {
			continue
		}

		// One of the initial registers.
		if strings.HasPrefix(row, "Register") {
			reRegister := regexp.MustCompile(`[A|B|C]`)
			reValue := regexp.MustCompile(`\d+`)
			register := reRegister.FindString(row)
			value := reValue.FindString(row)
			v, _ := strconv.Atoi(value)
			switch register {
			case "A":
				a = v
			case "B":
				b = v
			case "C":
				c = v
			}
		}

		// The program code itself.
		if strings.HasPrefix(row, "Program") {
			reProgram := regexp.MustCompile(`(\d+,?)+`)
			program := reProgram.FindString(row)
			p := strings.Split(program, ",")
			code = make([]int, len(p))
			for i, value := range p {
				v, _ := strconv.Atoi(value)
				code[i] = v
			}
		}
	}

	return a, b, c, code
}
