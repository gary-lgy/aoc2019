package main

import (
	"fmt"
	"io"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["13a"] = solve13a
	solvers["13b"] = solve13b
}

func solve13a(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	blockCount := 0
	output, _, err := intcode.RunSingleInstance(program, nil)
	for i, out := range output {
		if i%3 == 2 && out == 2 {
			blockCount++
		}
	}
	return fmt.Sprint(blockCount), nil
}

func solve13b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	program[0] = 2
	vm := intcode.NewVM(program)
	board, err := vm.Run(nil)
	if err != nil {
		return "", err
	}
	var ball, paddle, score int64
	outputHandler := func(board []int64) {
		for i := 0; i < len(board); i += 3 {
			x, y, z := board[i], board[i+1], board[i+2]
			switch {
			case z == 3:
				paddle = x
			case z == 4:
				ball = x
			case x == -1 && y == 0:
				score = z
			}
		}
	}
	outputHandler(board)
	for !vm.Stopped() {
		var move int64
		switch {
		case ball == paddle:
			move = 0
		case ball < paddle:
			move = -1
		case ball > paddle:
			move = 1
		}
		output, err := vm.Run([]int64{move})
		if err != nil {
			return "", err
		}
		outputHandler(output)
	}

	return fmt.Sprint(score), nil
}
