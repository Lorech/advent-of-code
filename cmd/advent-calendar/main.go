package main

import (
	"fmt"
	"lorech/advent-of-code-2024/pkg/file"
	"lorech/advent-of-code-2024/pkg/puzzles"
)

func main() {
	solved := 9
	for day := 1; day <= solved; day++ {
		data, error := file.ReadInfile(day)
		if error != nil {
			panic(error)
		}

		one, two, err := puzzles.Solve(day, data)
		if err != nil {
			fmt.Printf("Day %d: %v\n", day, err)
		} else {
			fmt.Printf("Day %d: %d, %d\n", day, one, two)
		}
	}
}
