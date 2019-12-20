package main

import (
	"fmt"
	"io"
	"math"

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

func amplifiersOutput(program []int64, phases []int) (int64, error) {
	var output int64 = 0
	for i := 0; i < len(phases); i++ {
		vm := intcode.NewVM(program)
		outputSlice, err := vm.Run([]int64{int64(phases[i]), output})
		if err != nil {
			return 0, err
		}
		output = outputSlice[0]
	}
	return output, nil
}

func chainedAmplifiersOutput(program []int64, phases []int) (int64, error) {
	l := len(phases)
	vms := make([]*intcode.VM, 0, l)
	for i := 0; i < l; i++ {
		vm := intcode.NewVM(program)
		_, err := vm.Run([]int64{int64(phases[i])})
		if err != nil {
			return 0, err
		}
		vms = append(vms, vm)
	}
	var output int64 = 0
	for {
		for i := 0; i < l; i++ {
			outputSlice, err := vms[i].Run([]int64{output})
			if err != nil {
				return 0, err
			}
			output = outputSlice[0]
		}
		if vms[l-1].Stopped() {
			break
		}
	}
	return output, nil
}

func maxAmplifiersOutput(input io.Reader, phases []int, outputFunc func([]int64, []int) (int64, error)) (int64, error) {
	program, err := intcode.ReadIntCode(input)
	if err != nil {
		return 0, err
	}
	permutations := permutations(phases)
	var max int64 = math.MinInt64
	for _, perm := range permutations {
		output, err := outputFunc(program, perm)
		if err != nil {
			return 0, err
		}
		if output > max {
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
