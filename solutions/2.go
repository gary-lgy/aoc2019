package aoc2019

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readProgram(input *os.File) []int {
	buf, err := ioutil.ReadAll(input)
	Check(err)
	tokens := strings.Split(strings.TrimSpace(string(buf)), ",")
	numbers := make([]int, len(tokens))
	for i := range tokens {
		tmp, err := strconv.ParseInt(tokens[i], 10, 32)
		Check(err)
		numbers[i] = int(tmp)
	}
	return numbers
}

func runProgram(n []int) int {
	numbers := make([]int, len(n))
	copy(numbers, n)
	pc := 0
	for pc < len(numbers) {
		opcode := numbers[pc]
		if opcode == 99 {
			break
		}

		op1, op2, dest := numbers[pc+1], numbers[pc+2], numbers[pc+3]
		switch opcode {
		case 1:
			numbers[dest] = numbers[op1] + numbers[op2]
		case 2:
			numbers[dest] = numbers[op1] * numbers[op2]
		default:
			panic("Unknown opcode")
		}

		pc += 4
	}

	return numbers[0]
}

func Solve2a(input *os.File) {
	numbers := readProgram(input)
	numbers[1], numbers[2] = 12, 2

	fmt.Println(runProgram(numbers))
}

func Solve2b(input *os.File) {
	numbers := readProgram(input)
	target := 19690720

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			numbers[1], numbers[2] = i, j
			if runProgram(numbers) == target {
				fmt.Println(100 * i + j)
				return
			}
		}
	}

	panic("Cannot find answer")
}
