# Advent of Code

A centralized helper and solution provider for Advent of Code puzzles.

## About

During my first run through Advent of Code in 2024, I immediately loved the challenge, and knew I would like to do more years and/or create solutions in multiple programming languages. Building off my solutions, this repository has since reemerged to provide me with the opportunity to pursue that goal, while also offering some extra goodies along the way:

- Randomized puzzle suggestions for use outside the advent
- Automatic bootstrapping of new puzzles (TODO)
- Centralized runner for comparing between solutions and languages (TODO)

## CLI

The root of the repository acts as an individual Go module, implementing a CLI tool for interfacing with Advent of Code itself. Specific examples of use cases will be demonstrated in the following sections, but remember - you can always run help, both on the root command, and subcommands to potentially find something more specific to your use case:

```sh
aoc --help
```

### Opening puzzles

During the advent, you can open today's puzzle:

```sh
aoc open
```

Running the same command outside the advent will open the nearest match to "today" as if it was the advent - the closest advent day's puzzle of the previous year's advent will be opened. But specific puzzles can also be opened, e.g., Day 14 of 2017:

```sh
aoc open -d 14 -y 2017
```

Want a challenge? Pick a random puzzle:

```sh
aoc random
```

Or, if you want _some_ structure, clamp your randomness to either a specific day, or a specific year, e.g., random puzzle of 2017:

```sh
aoc random -y 2017
```

### Bootstrapping puzzles

> [!WARNING]
> The following is not actually implemented, but I am including this to give an outline of my goals with the repository.

For quicker setup, especially during the advent, you can instantly bootstrap a new puzzle via [`templates`](https://github.com/Lorech/advent-of-code/tree/main/templates):

```sh
aoc setup --lang go
```

Templates are scoped to each language, where every file and folder within the directory will be transferred/modified, subject to the new puzzle's requirements. When creating a new template, some special keywords are defined, which you can use to fill out data scoped to the specific puzzle:

- `day` - numeric representation of the day this puzzle represents
- `year` - numeric representation of the year this puzzle represents

### Solving puzzles

> [!WARNING]
> The following is not actually implemented, but I am including this to give an outline of my goals with the repository.

Solutions are split by language inside [`solutions`](https://github.com/Lorech/advent-of-code/tree/main/solutions). Each language acts as a fully independent project, subjecting itself to the conventions applied to the specific language. Centralized solver executables are provided by each project, which can then be executed by the CLI:

Get the solution of a puzzle:

```sh
aoc solve -d 14 -y 17 --lang go
```

Get the test results of a puzzle:

```sh
aoc test -d 14 -y 17 --lang go
```

Get the benchmark results of a puzzle:

```sh
aoc bench -d 14 -y 17 --lang go
```

## Input files

The only input data stored within this repository is data created by myself for my own personal test cases. No official Advent of Code input data is provided as per the author's request.

To facilitate the use of multiple languages per puzzle, input data is centralized within the repository, see [`infiles`](https://github.com/Lorech/advent-of-code/tree/main/infiles) for details.

## Contributing

The repository contains my personal solutions, which are by no means the best - many of them may not even be good. Feel free to judge, but contributions to them will not be accepted. As per the main CLI functionality of the repository - feel free to provide suggestions, features, or bug fixes!
