package main

import (
	"fmt"
	"os"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solverMap["5a"] = solve5a
	solverMap["5b"] = solve5b
}

func solve5a(input *os.File) {
	vm := intcode.NewVM(intcode.ReadIntCode(input), []int{1})
	fmt.Println(vm.Run())
}

func solve5b(input *os.File) {
	vm := intcode.NewVM(intcode.ReadIntCode(input), []int{5})
	fmt.Println(vm.Run())
}
