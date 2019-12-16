package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solverMap["1a"] = solve1a
	solverMap["1b"] = solve1b
}

func fuelPartA(mass int) int {
	return mass/3 - 2
}

func solve1a(input *os.File) {
	scanner := bufio.NewScanner(input)

	var sum int = 0
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 32)
		aocutil.Check(err)
		sum += fuelPartA(int(mass))
	}
	fmt.Println(sum)
}

func fuelPartB(mass int) int {
	fuel := 0
	f := mass/3 - 2
	for f > 0 {
		fuel += f
		f = f/3 - 2
	}
	return fuel
}

func solve1b(input *os.File) {
	scanner := bufio.NewScanner(input)

	var sum int
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 32)
		aocutil.Check(err)
		sum += fuelPartB(int(mass))
	}
	fmt.Println(sum)
}
