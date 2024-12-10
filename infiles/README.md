# Input data

As per the author's request, no input data is stored in this repository:

> Please don't \[copy or redistribute parts of Advent of Code\]. Advent of Code is free to use, not free to copy. If you're posting a code repository somewhere, please don't include parts of Advent of Code like the puzzle text or your inputs. If you're making a website, please don't make it look like Advent of Code or name it something similar.

The files within this directory are files that are used for testing file reading from this directory, and do not contain any actual input data.

## Getting input data

To obtain input data for puzzles, you will need to sign up for [Advent of Code](https://adventofcode.com/), and obtain your personal input data from the website.

You can also create test files using the provided examples in each puzzle's description, which my unit tests have been created to compare against.

## Using input data

For the actual puzzles, save the your input file in this directory with the filename `{day}.txt`, where `{day}` is the day of the puzzle (e.g. `1.txt` for Day 1).

For the example puzzles, save the example data into a text file in this directory with the filename `{day}_test.txt`, where `{day}` is the day of the puzzle (e.g. `1_test.txt` for the example data for Day 1).

My file reader also supports adding variations of tests for each day, such as the additional examples provided on the daily puzzle or custom ones, but I intend on not adding actual unit tests for them because they become hard to maintain due to not being able to add them to VCS. If you spot one of these lingering, consider it a bug!
