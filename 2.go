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
	ic, oc := make(chan int), make(chan int)
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	vm := intcode.NewVM(program, ic, oc)
	vm.SetMemory(1, 12)
	vm.SetMemory(2, 2)
	go vm.Run()
	for o := range oc {
		intcode.LogOutput(o)
	}
	return fmt.Sprint(vm.ExitCode()), nil
}

func solve2b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	target := 19690720

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			ic, oc := make(chan int), make(chan int)
			vm := intcode.NewVM(program, ic, oc)
			vm.SetMemory(1, i)
			vm.SetMemory(2, j)
			go vm.Run()
			for o := range oc {
				intcode.LogOutput(o)
			}
			if vm.ExitCode() == target {
				return fmt.Sprint(100*i + j), nil
			}
		}
	}

	return "", fmt.Errorf("cannot find answer")
}
