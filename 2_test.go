package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gary-lgy/aoc2019/intcode"
	. "github.com/gary-lgy/aoc2019/testutil"
)

func TestVmReturnValue(t *testing.T) {
	tc := []VMTC{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{}, 3500, []int{}},
		{[]int{1, 0, 0, 0, 99}, []int{}, 2, []int{}},
		{[]int{2, 3, 0, 3, 99}, []int{}, 2, []int{}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{}, 2, []int{}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{}, 30, []int{}},
	}
	input, err := os.Open(filepath.Join("input", "2.txt"))
	require.NoError(t, err)
	defer input.Close()
	c1, err := intcode.ReadIntCode(input)
	require.NoError(t, err)
	c2 := make([]int, len(c1))
	copy(c2, c1)
	c1[1], c1[2] = 12, 2
	c2[1], c2[2] = 80, 18
	tc = append(tc,
		VMTC{Program: c1, Input: []int{}, ExpectedReturnValue: 3166704, ExpectedOutput: []int{}},
		VMTC{Program: c2, Input: []int{}, ExpectedReturnValue: 19690720, ExpectedOutput: []int{}})

	IntcodeVMTest(t, tc)
}
