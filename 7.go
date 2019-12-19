package main

import (
	"fmt"
	"io"
	"math"
	"sync"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solvers["7a"] = solve7a
	solvers["7b"] = solve7b
}

func permutationsHelper(numbers, current []int, results [][]int, available []bool) [][]int {
	if len(current) == len(numbers) {
		return append(results, current)
	}

	for i, ok := range available {
		if !ok {
			continue
		}
		available[i] = false
		next := make([]int, len(current), len(current)+1)
		copy(next, current)
		next = append(next, numbers[i])
		results = permutationsHelper(numbers, next, results, available)
		available[i] = true
	}
	return results
}

func permutations(numbers []int) [][]int {
	return permutationsHelper(numbers, []int{}, [][]int{}, []bool{true, true, true, true, true})
}

func amplifiersOutput(program []int64, phases []int) int64 {
	var output int64 = 0
	for i := 0; i < len(phases); i++ {
		ic, oc := make(chan int64), make(chan int64)
		vm := intcode.NewVM(program, ic, oc)
		go vm.Run()
		ic <- int64(phases[i])
		ic <- output
		output = <-oc
	}
	return output
}

func chainedAmplifiersOutput(program []int64, phases []int) int64 {
	l := len(phases)
	channels := make([]chan int64, 0, l)
	for i := 0; i < l; i++ {
		channels = append(channels, make(chan int64, 10))
	}
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		ic, oc := channels[i], channels[(i+1)%l]
		vm := intcode.NewVM(program, ic, oc)
		wg.Add(1)
		go vm.RunWithWG(&wg)
		ic <- int64(phases[i])
	}
	channels[0] <- 0
	wg.Wait()
	return <-channels[0]
}

func maxAmplifiersOutput(input io.Reader, phases []int, outputFunc func([]int64, []int) int64) (int64, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return 0, err
	}
	permutations := permutations(phases)
	var max int64 = math.MinInt64
	for _, perm := range permutations {
		if output := outputFunc(program, perm); output > max {
			max = output
		}
	}
	return max, nil
}

func solve7a(input io.Reader) (string, error) {
	answer, err := maxAmplifiersOutput(input, []int{0, 1, 2, 3, 4}, amplifiersOutput)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(answer), nil
}

func solve7b(input io.Reader) (string, error) {
	answer, err := maxAmplifiersOutput(input, []int{5, 6, 7, 8, 9}, chainedAmplifiersOutput)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(answer), nil
}
