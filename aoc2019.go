package main

import (
	"fmt"
	"os"

	aoc2019 "github.com/gary-lgy/aoc2019/solutions"
)

const usage = "Usage: aoc2019 PUZZLE [INPUT]"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	puzzle := os.Args[1]
	var solver func(*os.File)
	switch puzzle {
	case "1a":
		solver = aoc2019.Solve1a
	case "1b":
		solver = aoc2019.Solve1b
	default:
		fmt.Fprintln(os.Stderr, "Not implemented yet.")
		os.Exit(2)
	}

	var input string
	if len(os.Args) == 2 {
		input = puzzle
	} else {
		input = os.Args[2]
	}
	filename := "input/" + input
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file %q: %v\n", filename, err)
		os.Exit(3)
	}
	defer file.Close()

	solver(file)
}
