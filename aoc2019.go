package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	aoc2019 "github.com/gary-lgy/aoc2019/solutions"
)

type aocSolver func(*os.File)

const usage = "Usage: aoc2019 PUZZLE [INPUT]"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	puzzle := os.Args[1]
	solver, err := chooseSolver(puzzle)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	var input string
	if len(os.Args) == 2 {
		input = puzzle
	} else {
		input = os.Args[2]
	}
	filename := filepath.Join("input", input)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file %q: %v\n", filename, err)
		os.Exit(3)
	}
	defer file.Close()

	solver(file)
}

func chooseSolver(puzzle string) (aocSolver, error) {
	switch puzzle {
	case "1a":
		return aoc2019.Solve1a, nil
	case "1b":
		return aoc2019.Solve1b, nil
	case "2a":
		return aoc2019.Solve2a, nil
	case "2b":
		return aoc2019.Solve2b, nil
	case "3a":
		return aoc2019.Solve3a, nil
	case "3b":
		return aoc2019.Solve3b, nil
	case "4a":
		return aoc2019.Solve4a, nil
	case "4b":
		return aoc2019.Solve4b, nil
	default:
		return nil, errors.New("Not implemented yet.")
	}
}
