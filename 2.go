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
	vm := intcode.NewVM(intcode.ReadIntCode(input), []int{})
	vm.SetMemory(1, 12)
	vm.SetMemory(2, 2)
	fmt.Println(vm.Run())
}

func solve2b(input *os.File) {
	program := intcode.ReadIntCode(input)
	target := 19690720

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			vm := intcode.NewVM(program, []int{})
			vm.SetMemory(1, i)
			vm.SetMemory(2, j)
			if vm.Run() == target {
				fmt.Println(100*i + j)
				return
			}
		}
	}

	panic("Cannot find answer")
}
