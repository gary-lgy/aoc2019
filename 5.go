package main

import (
	"os"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solverMap["5a"] = solve5a
	solverMap["5b"] = solve5b
}

func solve5a(input *os.File) {
	ic, oc := make(chan int), make(chan int)
	vm := intcode.NewVM(intcode.ReadIntCode(input), ic, oc)
	go vm.Run()
	ic <- 1
	for output := range oc {
		intcode.LogOutput(output)
	}
}

func solve5b(input *os.File) {
	ic, oc := make(chan int), make(chan int)
	vm := intcode.NewVM(intcode.ReadIntCode(input), ic, oc)
	go vm.Run()
	ic <- 5
	for output := range oc {
		intcode.LogOutput(output)
	}
}
