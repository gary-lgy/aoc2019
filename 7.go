package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"

	"github.com/gary-lgy/aoc2019/intcode"
)

func init() {
	solverMap["7a"] = solve7a
	solverMap["7b"] = solve7b
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

func amplifiersOutput(program []int, phases []int) int {
	output := 0
	for i := 0; i < len(phases); i++ {
		ic, oc := make(chan int), make(chan int)
		vm := intcode.NewVM(program, ic, oc)
		go vm.Run()
		ic <- phases[i]
		ic <- output
		output = <-oc
	}
	return output
}

func chainedAmplifiersOutput(program []int, phases []int) int {
	l := len(phases)
	channels := make([]chan int, 0, l)
	for i := 0; i < l; i++ {
		channels = append(channels, make(chan int, 10))
	}
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		ic, oc := channels[i], channels[(i+1)%l]
		vm := intcode.NewVM(program, ic, oc)
		wg.Add(1)
		go vm.RunWithWG(&wg)
		ic <- phases[i]
	}
	channels[0] <- 0
	wg.Wait()
	return <-channels[0]
}

func maxAmplifiersOutput(input io.Reader, phases []int, outputFunc func([]int, []int) int) int {
	program := intcode.ReadIntCode(input)
	permutations := permutations(phases)
	max := math.MinInt32
	for _, perm := range permutations {
		if output := outputFunc(program, perm); output > max {
			max = output
		}
	}
	return max
}

func solve7a(input *os.File) {
	fmt.Println(maxAmplifiersOutput(input, []int{0, 1, 2, 3, 4}, amplifiersOutput))
}

func solve7b(input *os.File) {
	fmt.Println(maxAmplifiersOutput(input, []int{5, 6, 7, 8, 9}, chainedAmplifiersOutput))
}
