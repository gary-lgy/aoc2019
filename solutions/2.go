package aoc2019

import (
	"fmt"
	"os"
)

func Solve2a(input *os.File) {
	Ensure(RunProgram([]int{1,9,10,3,2,3,11,0,99,30,40,50}) == 3500)
	Ensure(RunProgram([]int{1,0,0,0,99}) == 2)
	Ensure(RunProgram([]int{2,3,0,3,99}) == 2)
	Ensure(RunProgram([]int{2,4,4,5,99,0}) == 2)
	Ensure(RunProgram([]int{1,1,1,4,99,5,6,0,99}) == 30)

	numbers := ReadProgram(input)
	numbers[1], numbers[2] = 12, 2
	fmt.Println(RunProgram(numbers))
}

func Solve2b(input *os.File) {
	numbers := ReadProgram(input)
	target := 19690720

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			numbers[1], numbers[2] = i, j
			if RunProgram(numbers) == target {
				fmt.Println(100*i + j)
				return
			}
		}
	}

	panic("Cannot find answer")
}
