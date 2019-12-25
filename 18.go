package main

import (
	"bufio"
	"fmt"
	"io"
	"unicode"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solvers["18a"] = solve18a
	solvers["18b"] = solve18b
}

func readMaze(input io.Reader) ([][]uint8, aocutil.IntPair) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)

	var (
		maze     [][]uint8
		line     []uint8
		entrance aocutil.IntPair
	)
	for scanner.Scan() {
		ch := scanner.Text()[0]
		switch ch {
		case '\n':
			maze = append(maze, line)
			line = []uint8{}
		case '@':
			entrance.X = len(line)
			entrance.Y = len(maze)
			line = append(line, '.')
		default:
			line = append(line, ch)
		}

	}
	return maze, entrance
}

func bfsDay18A(maze [][]uint8, entranceX, entranceY int) (int, error) {
	type point struct {
		X, Y int
		keys aocutil.BitSet
	}
	const allKeysCollected = 67108863                   // b00000011111111111111111111111111
	dx, dy := [4]int{0, 0, -1, 1}, [4]int{+1, -1, 0, 0} // ^, v, <, >

	origin := point{X: entranceX, Y: entranceY, keys: 0}
	var queue = []point{origin}
	steps := make(map[point]int)
	steps[origin] = 0
	for len(queue) > 0 {
		current := queue[0]
		if current.keys == allKeysCollected {
			return steps[current], nil
		}
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nextX, nextY := current.X+dx[i], current.Y+dy[i]
			if !(nextX >= 0 && nextX < len(maze[0]) && nextY >= 0 && nextY < len(maze)) {
				continue
			}
			neighbour := maze[nextY][nextX]
			if neighbour == '#' ||
				(unicode.IsUpper(rune(neighbour)) && !current.keys.Has(int(neighbour-'A'))) {
				continue
			}

			nextKeys := current.keys
			if unicode.IsLower(rune(neighbour)) {
				nextKeys.Set(int(neighbour - 'a'))
			}
			next := point{X: nextX, Y: nextY, keys: nextKeys}
			if _, exists := steps[next]; exists {
				continue
			}
			queue = append(queue, next)
			steps[next] = steps[current] + 1

		}
	}

	return 0, fmt.Errorf("failed to calculate the number of steps")
}

func solve18a(input io.Reader) (string, error) {
	maze, entrance := readMaze(input)
	steps, err := bfsDay18A(maze, entrance.X, entrance.Y)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(steps), nil
}
