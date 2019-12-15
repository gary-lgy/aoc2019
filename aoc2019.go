package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const usage = "Usage: aoc2019 PUZZLE [INPUT]"

type aocSolver func(*os.File)

var solverMap = make(map[string]aocSolver)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	puzzle := os.Args[1]
	solver, found := solverMap[puzzle]
	if !found {
		fmt.Fprintln(os.Stderr, errors.New("Not implemented yet."))
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
