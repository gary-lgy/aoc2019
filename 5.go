package main

import (
	"fmt"
	"os"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solverMap["5a"] = Solve5a
	solverMap["5b"] = Solve5b
}

func Solve5a(input *os.File) {
	vm := intcode.NewVm(intcode.ReadIntCode(input), []int{1})
	fmt.Println(vm.Run())
}

func Solve5b(input *os.File) {
	vm := intcode.NewVm(intcode.ReadIntCode(input), []int{5})
	fmt.Println(vm.Run())
}
