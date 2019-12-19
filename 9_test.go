package main

import (
	"path/filepath"
	"testing"

	"github.com/gary-lgy/aoc2019/testutil"
)

func TestDay9PartA(t *testing.T) {
	tc := []testutil.VMTC{
		{[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			nil, 109,
			[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
		{[]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, nil, 1102, []int64{1219070632396864}},
		{[]int64{104, 1125899906842624, 99}, nil, 104, []int64{1125899906842624}},
	}
	program := testutil.ReadIntcodeFromFile(t, filepath.Join("input", "9.txt"))
	tc = append(tc, testutil.VMTC{Program: program, Input: []int64{1}, ExpectedReturnValue: 1102, ExpectedOutput: []int64{3906448201}})

	testutil.IntcodeVMTest(t, tc)
}

func TestDay9PartB(t *testing.T) {
	program := testutil.ReadIntcodeFromFile(t, filepath.Join("input", "9.txt"))
	tc := testutil.VMTC{Program: program, Input: []int64{2}, ExpectedReturnValue: 1102, ExpectedOutput: []int64{59785}}
	testutil.IntcodeVMTest(t, []testutil.VMTC{tc})
}
