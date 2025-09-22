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

Solves depend on the personalized input files generated for each user. These files must be stored within the `infiles` directory at the root of the repository, named `{year}/{day}.txt`, e.g., `2024/09.txt`.

Solve every puzzle with a solution:

```bash
go run cmd/solve/main.go
```

Solve a specific day's puzzle:

```bash
go run cmd/solve/main.go -day 9
```

### Running tests

Tests depend on the example solutions part of each puzzle's description. These files must be stored within the `infiles` directory at the root of the repository, named `{year}/{day}_test.txt`, e.g., `2024/09_test.txt`.

Test every puzzle with solutions:

```bash
go test pkg/puzzles
```

Test a specific day's puzzle:

```bash
go test pkg/puzzles -run Nine
```

### Running benchmarks

Benchmarks depend on the personalized input files generated for each user. These files must be stored within the `infiles` directory at the root of the repository, named `{year}/{day}.txt`, e.g., `2024/09.txt`.

Benchmark every puzzle with a solution:

```bash
go test pkg/puzzles -bench .
```

Benchmark a specific day's puzzle:

```bash
go test pkg/puzzles bench Nine
```
