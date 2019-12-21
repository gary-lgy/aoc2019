package main

import (
	"fmt"
	"io"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["15a"] = solve15a
	solvers["15b"] = solve15b
}

func nextPos(current aocutil.IntPair, direction int) (aocutil.IntPair, error) {
	// north (1), south (2), west (3), and east (4)
	switch direction {
	case 1:
		return aocutil.IntPair{X: current.X, Y: current.Y + 1}, nil
	case 2:
		return aocutil.IntPair{X: current.X, Y: current.Y - 1}, nil
	case 3:
		return aocutil.IntPair{X: current.X - 1, Y: current.Y}, nil
	case 4:
		return aocutil.IntPair{X: current.X + 1, Y: current.Y}, nil
	default:
		return aocutil.IntPair{}, fmt.Errorf("unknown direction")
	}
}

func reverseDirection(direction int) int {
	// 1 -> 2, 2 -> 1, 3 -> 4, 4 -> 3
	if direction%2 == 0 {
		return direction - 1
	} else {
		return direction + 1
	}
}

// Explore the entire map and find the location of and distance from the oxygen system
func findOxygenSystem(vm *intcode.VM, pos aocutil.IntPair, environment map[aocutil.IntPair]int, steps int, systemDist *int, systemPos *aocutil.IntPair) error {
	for i := 1; i <= 4; i++ {
		next, err := nextPos(pos, i)
		if err != nil {
			return err
		}
		if _, visited := environment[next]; visited {
			continue
		}
		output, err := vm.Run([]int64{int64(i)})
		if err != nil {
			return err
		}
		switch output[0] {
		case 0: // hit a wall
			environment[next] = 0
		case 2: // found oxygen system
			*systemDist, *systemPos = steps+1, next
			fallthrough
		case 1: // moved
			environment[next] = int(output[0])
			if err := findOxygenSystem(vm, next, environment, steps+1, systemDist, systemPos); err != nil {
				return err
			}
			_, err = vm.Run([]int64{int64(reverseDirection(i))})
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown output from VM")
		}
	}
	return nil
}

func solve15a(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	vm := intcode.NewVM(program)
	environment := make(map[aocutil.IntPair]int)
	environment[aocutil.IntPair{X: 0, Y: 0}] = 1
	var (
		systemDist int
		systemPos  aocutil.IntPair
	)
	if err := findOxygenSystem(vm, aocutil.IntPair{X: 0, Y: 0}, environment, 0, &systemDist, &systemPos); err != nil {
		return "", err
	}
	return fmt.Sprint(systemDist), nil
}

func bfsFillWithOxygen(environment map[aocutil.IntPair]int, systemPos aocutil.IntPair) (int, error) {
	type s struct {
		pos  aocutil.IntPair
		dist int
	}
	var queue []s
	queue = append(queue, s{systemPos, 0})
	filled := make(map[aocutil.IntPair]bool)
	filled[systemPos] = true
	time := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		time = aocutil.MaxInt(time, current.dist)
		for i := 1; i <= 4; i++ {
			next, err := nextPos(current.pos, i)
			if err != nil {
				return 0, err
			}
			if status, charted := environment[next]; !charted || status == 0 {
				continue
			}
			if _, f := filled[next]; !f {
				filled[next] = true
				queue = append(queue, s{next, current.dist + 1})
			}
		}
	}
	return time, nil
}

func solve15b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	vm := intcode.NewVM(program)
	environment := make(map[aocutil.IntPair]int)
	environment[aocutil.IntPair{X: 0, Y: 0}] = 1
	var (
		systemDist int
		systemPos  aocutil.IntPair
	)
	if err := findOxygenSystem(vm, aocutil.IntPair{X: 0, Y: 0}, environment, 0, &systemDist, &systemPos); err != nil {
		return "", err
	}
	timeRequired, err := bfsFillWithOxygen(environment, systemPos)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(timeRequired), nil
}
