package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func init() {
	solvers["4a"] = solve4a
	solvers["4b"] = solve4b
}

func readRange(input io.Reader) (int, int, error) {
	buf, err := ioutil.ReadAll(input)
	if err != nil {
		return 0, 0, err
	}
	data := strings.Split(strings.TrimSpace(string(buf)), "-")
	l, e1 := strconv.ParseInt(data[0], 10, 32)
	if e1 != nil {
		return 0, 0, e1
	}
	h, e2 := strconv.ParseInt(data[1], 10, 32)
	if e2 != nil {
		return 0, 0, e2
	}
	return int(l), int(h), nil
}

func isPossiblePartA(password int) bool {
	repr := strconv.Itoa(password)
	haveDouble := false
	for i := 1; i < len(repr); i++ {
		if repr[i-1] > repr[i] {
			return false
		} else if repr[i-1] == repr[i] {
			haveDouble = true
		}
	}
	return haveDouble
}

func solve4aTheStupidWay(input io.Reader) (int, error) {
	low, high, err := readRange(input)
	if err != nil {
		return 0, err
	}
	ways := 0
	for i := low; i <= high; i++ {
		if isPossiblePartA(i) {
			ways++
		}
	}
	return ways, nil
}

func isPossiblePartB(password int) bool {
	repr := strconv.Itoa(password)
	repeated := 1
	haveDouble := false
	for i := 1; i < len(repr); i++ {
		if repr[i-1] > repr[i] {
			return false
		} else if repr[i-1] == repr[i] {
			repeated++
		} else {
			if repeated == 2 {
				haveDouble = true
			}
			repeated = 1
		}
	}
	return haveDouble || repeated == 2
}

func solve4bTheStupidWay(input io.Reader) (int, error) {
	low, high, err := readRange(input)
	if err != nil {
		return 0, err
	}
	ways := 0
	for i := low; i <= high; i++ {
		if isPossiblePartB(i) {
			ways++
		}
	}
	return ways, nil
}

func solve4a(input io.Reader) (string, error) {
	answer, err := solve4aTheStupidWay(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(answer), nil
}

func solve4b(input io.Reader) (string, error) {
	answer, err := solve4bTheStupidWay(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(answer), nil
}
