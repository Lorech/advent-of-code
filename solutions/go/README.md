# Advent of Code 2024 

This is my personal hub for the Advent of Code 2024, which I am solving using [Go](https://golang.org/).

## Structure

Primarily, my repository consists of three parts:

- A main command that runs the possible solutions;
- Unit tests, validating the solutions on the author's provided examples;
- Benchmarks, measuring the performance of the solutions using real input data.

Over time, I plan on expanding the repository further, using this as a playground for learning Go, and possibly other languages.

## Usage

### Solving puzzles

To solve puzzles for every day that I have a solution for:

```bash
go run cmd/solve/main.go
```

To solve puzzles for a specific day only (e.g. day 9):

```bash'
go run cmd/solve/main.go -day 9
```

### Running tests

To run all tests:

```bash
go test pkg/puzzles
```

To run tests for a specific day (e.g. day 9):

```bash
go test pkg/puzzles -run Nine
```

### Running benchmarks

To run all benchmarks:

```bash
go test pkg/puzzles -bench .
```

To run benchmarks for a specific day (e.g. day 9):

```bash
go test pkg/puzzles bench Nine
```

## Input data

As per the author's request, no input data is stored in this repository. All data used for puzzles must be obtained from the website directly.

Two types of input files are used and commited:

- `infiles/{day}.txt` for the puzzle input;
- `infiles/{day}_test.txt` for the example input.

For more details about obtaining and storing input data, see [the infile README](infiles/README.md).
