package puzzles

import (
	"slices"
	"strconv"
	"strings"
)

type block struct {
	val int // The value of the block.
	len int // The times this value repeats at this position.
}

// Day 9: Disk Fragmenter
// https://adventofcode.com/2024/day/9
func dayNine(input string) (int, int) {
	return d9p1(input), 0
}

// Completes the first half of the puzzle for day 9.
func d9p1(input string) int {
	disk := parseDisk(input)
	checksum := 0

nextGap:
	for i := 0; i < len(disk); i++ {
		// Outer loop only deals with gaps.
		if disk[i].val != -1 {
			continue
		}

		for j := len(disk) - 1; j > i; j-- {
			// Inner loop only deals with blocks.
			if disk[j].val == -1 {
				disk = disk[:j]
				continue
			}

			if disk[i].len == disk[j].len {
				disk[i].val = disk[j].val
				disk = disk[:j]
				goto nextGap
			} else if disk[j].len > disk[i].len {
				disk[i].val = disk[j].val
				disk[j].len -= disk[i].len
				goto nextGap
			} else if disk[i].len > disk[j].len {
				l := disk[i].len
				disk[i].val = disk[j].val
				disk[i].len = disk[j].len
				disk = slices.Insert(disk, i+1, block{-1, l - disk[j].len})
				disk = disk[:j+1]
				break
			}
		}
	}

	p := 0
	for _, b := range disk {
		// Ignore any remaining free space.
		if b.val == -1 {
			continue
		}

		for j := range b.len {
			checksum += (p + j) * b.val
		}
		p += b.len
	}

	return checksum
}
	}

	return checksum
}

// Parses the input data, converting it to an organized structure of the disk drive.
func parseDisk(input string) []block {
	compression, _ := strings.CutSuffix(input, "\n")
	disk := make([]block, 0)

	for i := 0; i < len(compression); i += 2 {
		f := i / 2
		l, _ := strconv.Atoi(string(compression[i]))
		file := block{f, l}

		// Index the file into a slice of blocks.
		disk = append(disk, file)

		// Every file except for the last one also has free space after it.
		if i != len(compression)-1 {
			s, _ := strconv.Atoi(string(compression[i+1]))
			space := block{-1, s}
			disk = append(disk, space)
		}
	}

	return disk
}
