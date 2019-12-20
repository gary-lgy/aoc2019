package main

import (
	"fmt"
	"io"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["2a"] = solve2a
	solvers["2b"] = solve2b
}

func solve2a(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	program[1], program[2] = 12, 2
	output, exitCode, err := intcode.RunSingleInstance(program, nil)
	if err != nil {
		return "", err
	}
	for _, out := range output {
		intcode.LogOutput(out)
	}
	return fmt.Sprint(exitCode), nil
}

func solve2b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	var target int64 = 19690720
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			program[1], program[2] = int64(i), int64(j)
			output, exitCode, err := intcode.RunSingleInstance(program, nil)
			if err != nil {
				return "", err
			}
			for _, out := range output {
				intcode.LogOutput(out)
			}
			if exitCode == target {
				return fmt.Sprint(100*i + j), nil
			}
		}
	}

	return "", fmt.Errorf("cannot find answer")
}
