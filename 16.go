package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solvers["16a"] = solve16a
	solvers["16b"] = solve16b
}

func readNumbers(input string) ([]int8, error) {
	var numbers []int8
	for i := range input {
		num, err := strconv.ParseInt(input[i:i+1], 10, 5)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int8(num))
	}
	return numbers, nil
}

func fft(numbers []int8) []int8 {
	l := len(numbers)
	output := make([]int8, l)
	for i := range numbers {
		multiplier := 1
		newNum := 0
		for j := i; j < l; {
			for k := 0; j+k < l && k < i+1; k++ {
				newNum += int(numbers[j+k]) * multiplier
			}
			j += 2 * (i + 1)
			multiplier *= -1
		}
		output[i] = int8(aocutil.AbsInt(newNum) % 10)
	}

	return output
}

func solve16a(input io.Reader) (string, error) {
	raw, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}
	str := strings.TrimSpace(string(raw))
	numbers, err := readNumbers(str)
	if err != nil {
		return "", err
	}
	phases := 100
	for i := 0; i < phases; i++ {
		numbers = fft(numbers)
	}
	return fmt.Sprint(numbers[8]), nil
}

func neededNumbers(numbers []int8, offset int) []int8 {
	l := len(numbers)*10000 - offset
	needed := make([]int8, l, l)
	for i := 0; i < l; i++ {
		needed[i] = numbers[(offset+i)%len(numbers)]
	}
	return needed
}

func solve16b(input io.Reader) (string, error) {
	raw, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}
	str := strings.TrimSpace(string(raw))
	numbers, err := readNumbers(str)
	if err != nil {
		return "", err
	}
	offset, err := strconv.ParseInt(str[:7], 10, 32)
	if err != nil {
		return "", err
	}
	if int(offset) < len(numbers)*10000/2 {
		return "", fmt.Errorf("algorithm failed: message offset is less than half of input length")
	}
	numbers = neededNumbers(numbers, int(offset))

	phases := 100
	for i := 0; i < phases; i++ {
		l := len(numbers)
		next := make([]int8, l)
		next[l-1] = numbers[l-1]
		for j := l - 2; j >= 0; j-- {
			next[j] = (next[j+1] + numbers[j]) % 10
		}
		numbers = next
	}
	return fmt.Sprint(numbers[:8]), nil
}
