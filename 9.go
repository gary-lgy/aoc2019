package main

import (
	"io"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["9a"] = solve9a
	solvers["9b"] = solve9b
}

func solve9a(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	output, _, err := intcode.RunSingleInstance(program, []int64{1})
	if err != nil {
		return "", err
	}
	for _, out := range output {
		intcode.LogOutput(out)
	}
	return "", nil
}

func solve9b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	output, _, err := intcode.RunSingleInstance(program, []int64{2})
	if err != nil {
		return "", err
	}
	for _, out := range output {
		intcode.LogOutput(out)
	}
	return "", nil
}

