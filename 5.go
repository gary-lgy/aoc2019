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
	ic, oc := make(chan int), make(chan int)
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	vm := intcode.NewVM(program, ic, oc)
	go vm.Run()
	ic <- 1
	for output := range oc {
		intcode.LogOutput(output)
	}
	return "", nil
}

func solve5b(input io.Reader) (string, error) {
	ic, oc := make(chan int), make(chan int)
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	vm := intcode.NewVM(program, ic, oc)
	go vm.Run()
	ic <- 5
	for output := range oc {
		intcode.LogOutput(output)
	}
	return "", nil
}
