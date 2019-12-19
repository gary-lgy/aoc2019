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
	intcode.RunSingleInstance(program, []int64{1}, intcode.LogOutput)
	return "", nil
}

func solve9b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	intcode.RunSingleInstance(program, []int64{2}, intcode.LogOutput)
	return "", nil
}

