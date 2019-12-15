package main

import (
	"fmt"
	"os"

	"github.com/gary-lgy/aoc2019/intcode"
)

func Solve5a(input *os.File) {
	vm := intcode.NewVm(intcode.ReadIntCode(input), []int{1})
	fmt.Println(vm.Run())
}

func Solve5b(input *os.File) {
	vm := intcode.NewVm(intcode.ReadIntCode(input), []int{5})
	fmt.Println(vm.Run())
}
