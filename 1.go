package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func init() {
	solvers["1a"] = solve1a
	solvers["1b"] = solve1b
}

func fuelPartA(mass int) int {
	return mass/3 - 2
}

func solve1a(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var sum = 0
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return "", err
		}
		sum += fuelPartA(int(mass))
	}
	return fmt.Sprint(sum), nil
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

func solve1b(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var sum int
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return "", err
		}
		sum += fuelPartB(int(mass))
	}
	return fmt.Sprint(sum), nil
}
