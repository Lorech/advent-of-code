package puzzles

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Day 17: Chronospacial Computer
// https://adventofcode.com/2024/day/17
func daySeventeen(input string) (interface{}, interface{}) {
	return d17p1(input), 0
}

// Completes the first half of the puzzle for day 17.
func d17p1(input string) string {
	a, b, c, program := parseProgram(input)
	return runProgram(program, &a, &b, &c)
}

// Runs the provided program, returning the string the program would output.
func runProgram(program []int, a *int, b *int, c *int) string {
	var o strings.Builder
	ip := 0

	for ip < len(program)-1 {
		opcode, literal := program[ip], program[ip+1]

		switch opcode {
		case 0: // ADV
			*a = *a / int(math.Pow(2, float64(comboOperand(literal, *a, *b, *c))))
		case 1: // BXL
			*b = *b ^ literal
		case 2: // BST
			*b = comboOperand(literal, *a, *b, *c) % 8
		case 3: // JNZ
			if *a != 0 {
				ip = literal
				continue // The pointer does not move if we jump.
			}
		case 4: // BXC
			*b = *b ^ *c
		case 5: // OUT
			v := comboOperand(literal, *a, *b, *c) % 8
			if o.Len() != 0 {
				o.WriteString(fmt.Sprintf(",%d", v))
			} else {
				o.WriteString(fmt.Sprintf("%d", v))
			}
		case 6: // BDV
			*b = *a / int(math.Pow(2, float64(comboOperand(literal, *a, *b, *c))))
		case 7: // CDV
			*c = *a / int(math.Pow(2, float64(comboOperand(literal, *a, *b, *c))))
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
