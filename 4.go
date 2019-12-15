package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	. "github.com/gary-lgy/aoc2019/aocutil"
)

func readRange(input *os.File) (low, high int) {
	buf, err := ioutil.ReadAll(input)
	Check(err)
	data := strings.Split(strings.TrimSpace(string(buf)), "-")
	l, e1 := strconv.ParseInt(data[0], 10, 32)
	Check(e1)
	h, e2 := strconv.ParseInt(data[1], 10, 32)
	Check(e2)
	low, high = int(l), int(h)
	return
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

func Solve4aTheStupidWay(input *os.File) {
	low, high := readRange(input)
	ways := 0
	for i := low; i <= high; i++ {
		if isPossiblePartA(i) {
			ways += 1
		}
	}
	fmt.Println(ways)
}

func isPossiblePartB(password int) bool {
	repr := strconv.Itoa(password)
	repeated := 1
	haveDouble := false
	for i := 1; i < len(repr); i++ {
		if repr[i-1] > repr[i] {
			return false
		} else if repr[i-1] == repr[i] {
			repeated += 1
		} else {
			if repeated == 2 {
				haveDouble = true
			}
			repeated = 1
		}
	}
	return haveDouble || repeated == 2
}

func Solve4bTheStupidWay(input *os.File) {
	low, high := readRange(input)
	ways := 0
	for i := low; i <= high; i++ {
		if isPossiblePartB(i) {
			ways += 1
		}
	}
	fmt.Println(ways)
}

func Solve4a(input *os.File) {
	Ensure(isPossiblePartA(111111))
	Ensure(!isPossiblePartA(223450))
	Ensure(!isPossiblePartA(123789))
	Solve4aTheStupidWay(input)
}

func Solve4b(input *os.File) {
	Ensure(isPossiblePartB(112233))
	Ensure(!isPossiblePartB(123444))
	Ensure(isPossiblePartB(111122))
	Solve4bTheStupidWay(input)
}
