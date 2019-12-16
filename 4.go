package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solverMap["4a"] = solve4a
	solverMap["4b"] = solve4b
}

func readRange(input *os.File) (low, high int) {
	buf, err := ioutil.ReadAll(input)
	aocutil.Check(err)
	data := strings.Split(strings.TrimSpace(string(buf)), "-")
	l, e1 := strconv.ParseInt(data[0], 10, 32)
	aocutil.Check(e1)
	h, e2 := strconv.ParseInt(data[1], 10, 32)
	aocutil.Check(e2)
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

func solve4aTheStupidWay(input *os.File) {
	low, high := readRange(input)
	ways := 0
	for i := low; i <= high; i++ {
		if isPossiblePartA(i) {
			ways++
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

func solve4bTheStupidWay(input *os.File) {
	low, high := readRange(input)
	ways := 0
	for i := low; i <= high; i++ {
		if isPossiblePartB(i) {
			ways++
		}
	}
	fmt.Println(ways)
}

func solve4a(input *os.File) {
	solve4aTheStupidWay(input)
}

func solve4b(input *os.File) {
	solve4bTheStupidWay(input)
}
