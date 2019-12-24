package main

import (
	"fmt"
	"io"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["17a"] = solve17a
	solvers["17b"] = solve17b
}

func parseCameraView(view []int64) [][]int8 {
	var env [][]int8
	var line []int8
	for i := range view {
		ch := int8(view[i])
		if ch == '\n' {
			if len(line) > 1 {
				env = append(env, line)
			}
			line = []int8{}
		} else {
			line = append(line, ch)
		}
	}
	return env
}

func isScaffold(point int8) bool {
	switch point {
	case '#', '^', 'v', '<', '>', 'X':
		return true
	default:
		return false
	}
}

func findIntersections(view [][]int8) []aocutil.IntPair {
	var intersections []aocutil.IntPair
	for i := 1; i < len(view)-1; i++ {
		for j := 1; j < len(view[i])-1; j++ {
			if isScaffold(view[i][j]) &&
				isScaffold(view[i-1][j]) &&
				isScaffold(view[i+1][j]) &&
				isScaffold(view[i][j-1]) &&
				isScaffold(view[i][j+1]) {
				intersections = append(intersections, aocutil.IntPair{X: i, Y: j})
			}
		}
	}
	return intersections
}

func solve17a(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	vm := intcode.NewVM(program)
	v, err := vm.Run([]int64{})
	view := parseCameraView(v)
	intersections := findIntersections(view)
	sum := 0
	for _, x := range intersections {
		sum += x.X * x.Y
	}
	return fmt.Sprint(sum), nil
}

func solve17b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	program[0] = 2
	vm := intcode.NewVM(program)
	// Hard-coded answer for my input -_-
	output, err := vm.Run([]int64{'A', ',', 'C', ',', 'A', ',', 'B', ',', 'C', ',', 'B', ',', 'C', ',', 'A', ',', 'B', ',', 'C', '\n',
		'L', ',', '1', '0', ',', 'L', ',', '6', ',', 'R', ',', '1', '0', '\n',
		'L', ',', '1', '0', ',', 'R', ',', '8', ',', 'R', ',', '8', ',', 'L', ',', '1', '0', '\n',
		'R', ',', '6', ',', 'R', ',', '8', ',', 'R', ',', '8', ',', 'L', ',', '6', ',', 'R', ',', '8', '\n',
		'n', '\n'})
	if err != nil {
		return "", err
	}
	return fmt.Sprint(output[len(output) - 1]), nil
}
