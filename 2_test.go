package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
	. "github.com/gary-lgy/aoc2019/testutil"
)

func TestVmReturnValue(t *testing.T) {
	tc := []VmTc{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{}, 3500, []int{}},
		{[]int{1, 0, 0, 0, 99}, []int{}, 2, []int{}},
		{[]int{2, 3, 0, 3, 99}, []int{}, 2, []int{}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{}, 2, []int{}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{}, 30, []int{}},
	}
	input, err := os.Open(filepath.Join("input", "2a"))
	defer input.Close()
	aocutil.Check(err)
	c1 := intcode.ReadIntCode(input)
	c2 := make([]int, len(c1))
	copy(c2, c1)
	c1[1], c1[2] = 12, 2
	c2[1], c2[2] = 80, 18
	tc = append(tc, VmTc{c1, []int{}, 3166704, []int{}}, VmTc{c2, []int{}, 19690720, []int{}})

	IntcodeVmTest(t, tc)
}
