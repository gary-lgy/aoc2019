package main

import (
	"io"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["5a"] = solve5a
	solvers["5b"] = solve5b
}

func solve5a(input io.Reader) (string, error) {
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

func solve5b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	output, _, err := intcode.RunSingleInstance(program, []int64{5})
	if err != nil {
		return "", err
	}
	for _, out := range output {
		intcode.LogOutput(out)
	}
	return "", nil
}
