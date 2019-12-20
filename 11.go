package main

import (
	"fmt"
	"io"
	"math"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["11a"] = solve11a
	solvers["11b"] = solve11b
}

func nextDirection(direction, turn int) (int, error) {
	switch turn {
	case 0: // left
		return (direction + 4 - 1) % 4, nil
	case 1: // right
		return (direction + 1) % 4, nil
	default:
		return 0, fmt.Errorf("unknown turn")
	}
}

func moveForward(pos aocutil.IntPair, direction int) (aocutil.IntPair, error) {
	switch direction {
	case 0: // up
		return aocutil.IntPair{X: pos.X, Y: pos.Y + 1}, nil
	case 1: // right
		return aocutil.IntPair{X: pos.X + 1, Y: pos.Y}, nil
	case 2: // down
		return aocutil.IntPair{X: pos.X, Y: pos.Y - 1}, nil
	case 3: // left
		return aocutil.IntPair{X: pos.X - 1, Y: pos.Y}, nil
	default:
		return aocutil.IntPair{}, fmt.Errorf("unknown direction")
	}
}

func paintHull(program []int64, initialColor int) (map[aocutil.IntPair]int, error) {
	ic, oc := make(chan int64), make(chan int64)
	vm := intcode.NewVM(program, ic, oc)
	go vm.Run()

	// 0: black; 1: white
	colors := make(map[aocutil.IntPair]int)
	// 0: up; 1: right, 2: down; 3: left
	direction := 0
	origin, pos := aocutil.IntPair{X: 0, Y: 0}, aocutil.IntPair{X: 0, Y: 0}
	// FIXME: this check is subject to race condition.
	// If the main coroutine reaches this check before the vm coroutine executes instruction 99, deadlock will occur
	for !vm.Stopped() {
		if color, exists := colors[pos]; exists {
			ic <- int64(color)
		} else if pos == origin {
			ic <- int64(initialColor)
		} else {
			ic <- 0
		}
		colorToPaint := <-oc
		colors[pos] = int(colorToPaint)
		turn := <-oc
		newDirection, err := nextDirection(direction, int(turn))
		if err != nil {
			return nil, err
		}
		direction = newDirection
		pos, err = moveForward(pos, direction)
		if err != nil {
			return nil, err
		}
	}
	return colors, nil
}

func solve11a(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	colors, err := paintHull(program, 0)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(len(colors)), nil
}

func solve11b(input io.Reader) (string, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return "", err
	}
	colors, err := paintHull(program, 1)
	if err != nil {
		return "", err
	}
	return drawShipID(colors), nil
}

func drawShipID(colors map[aocutil.IntPair]int) string {
	xMin, xMax, yMin, yMax := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
	for pos := range colors {
		xMin = aocutil.MinInt(xMin, pos.X)
		xMax = aocutil.MaxInt(xMax, pos.X)
		yMin = aocutil.MinInt(yMin, pos.Y)
		yMax = aocutil.MaxInt(yMax, pos.Y)
	}
	id := []byte{'\n'}
	for i := yMax; i >= yMin; i-- {
		for j := xMin; j <= xMax; j++ {
			if color, exists := colors[aocutil.IntPair{X: j, Y: i}]; exists && color == 1 {
				id = append(id, '#') // white
			} else {
				id = append(id, ' ') // black
			}
		}
		id = append(id, '\n')
	}
	return string(id)
}
