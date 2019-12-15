package main

import (
	"fmt"
	"os"

	"github.com/gary-lgy/aoc2019/intcode"
)

func Solve2a(input *os.File) {
	vm := intcode.NewVm(intcode.ReadIntCode(input), []int{})
	vm.SetMemory(1, 12)
	vm.SetMemory(2, 2)
	fmt.Println(vm.Run())
}

func Solve2b(input *os.File) {
	program := intcode.ReadIntCode(input)
	target := 19690720

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			vm := intcode.NewVm(program, []int{})
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
