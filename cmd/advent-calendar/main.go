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

	result := puzzles.DayOne(string(data))
	fmt.Println(result)
}
