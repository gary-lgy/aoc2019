package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const usage = "Usage: aoc2019 PUZZLE [INPUT]"

type aocSolver func(io.Reader) (string, error)

var solvers = make(map[string]aocSolver)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	puzzle := os.Args[1]
	solver, found := solvers[puzzle]
	if !found {
		fmt.Fprintln(os.Stderr, errors.New("not implemented yet"))
		os.Exit(2)
	}

	var input string
	if len(os.Args) == 2 {
		// If no filename is provided, infer it from puzzle name
		input = filepath.Join("input", strings.TrimRight(puzzle, "ab") + ".txt")
	} else {
		input = os.Args[2]
	}
	file, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file %q: %v\n", input, err)
		os.Exit(3)
	}
	defer file.Close()

	answer, err := solver(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to solve %s: %v\n", puzzle, err)
		os.Exit(4)
	}
	fmt.Println("Answer:", answer)
}
