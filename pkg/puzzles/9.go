package puzzles

import (
	"lorech/advent-of-code-2024/pkg/cslices"
	"strconv"
	"strings"
)

// Day 9: Disk Fragmenter
// https://adventofcode.com/2024/day/9
func dayNine(input string) (int, int) {
	return d9p1(input), 0
}

// Completes the first half of the puzzle for day 9.
func d9p1(input string) int {
	disk := parseDisk(input)
	checksum := 0

	for true {
		start, end := cslices.Appears(disk, -1)

		// There is no more free space to insert data into.
		if start == -1 {
			break
		}

		// There is no more data at the end of the disk.
		if end == len(disk) {
			disk = disk[:start+1]
		}

		length := end - start
		data := make([]int, length)

		i := 0
		l := 0
		for l < length {
			s := len(disk) - 1 - i
			i++

			// This empty block is larger than the data that can be inserted.
			if s < end {
				data = data[:l+1]
				break
			}

			// There is no data to get from this block.
			if disk[s] == -1 {
				continue
			}

			data[l] = disk[s]
			l++
		}

		for i, d := range data {
			disk[start+i] = d
		}

		disk = disk[:len(disk)-i]
	}

	for i, v := range disk {
		checksum += i * v
	}

	return checksum
}

// Parses the input data, converting it to an uncompressed disk drive.
func parseDisk(input string) []int {
	compression, _ := strings.CutSuffix(input, "\n")
	disk := make([]int, 0)

	for i := 0; i < len(compression); i += 2 {
		f := i / 2
		l, _ := strconv.Atoi(string(compression[i]))
		s := 0

		// Free space is only added for blocks before the final one.
		if i != len(compression)-1 {
			s, _ = strconv.Atoi(string(compression[i+1]))
		}

		// Map the file onto the disk.
		files := make([]int, l)
		for j := range files {
			files[j] = f
		}
		disk = append(disk, files...)

		// Map the free space onto the disk.
		space := make([]int, s)
		for j := range space {
			space[j] = -1
		}
		disk = append(disk, space...)
	}

	return disk
}
