package main

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/puzzles"
	"os"
)

func main() {
	data, error := os.ReadFile("infiles/1.txt")
	if error != nil {
		panic(error)
	}

	one, two := puzzles.DayOne(string(data))
	fmt.Printf("Day 1, part 1 (distance): %d, part 2 (similarity): %d\n", one, two)
}
