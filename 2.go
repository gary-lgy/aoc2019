package main

import (
	"fmt"
	"os"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solverMap["2a"] = solve2a
	solverMap["2b"] = solve2b
}

func solve2a(input *os.File) {
	ic, oc := make(chan int), make(chan int)
	vm := intcode.NewVM(intcode.ReadIntCode(input), ic, oc)
	vm.SetMemory(1, 12)
	vm.SetMemory(2, 2)
	go vm.Run()
	for o := range oc {
		intcode.LogOutput(o)
	}
	fmt.Println(vm.ExitCode())
}

func solve2b(input *os.File) {
	program := intcode.ReadIntCode(input)
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
				fmt.Println(100*i + j)
				return
			}
		}
	}

	panic("Cannot find answer")
}
