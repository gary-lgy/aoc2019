package aoc2019

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve1a(input *os.File) {
	scanner := bufio.NewScanner(input)

	var sum int64 = 0
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 32)
		Check(err)
		sum += mass/3 - 2
	}
	fmt.Println(sum)
}

func Solve1b(input *os.File) {
	scanner := bufio.NewScanner(input)

	var sum int64
	for scanner.Scan() {
		mass, err := strconv.ParseInt(scanner.Text(), 10, 32)
		Check(err)
		f := mass/3 - 2
		for f > 0 {
			sum += f
			f = f/3 - 2
		}
	}
	fmt.Println(sum)
}
