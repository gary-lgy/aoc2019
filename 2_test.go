package main

import (
	"path/filepath"
	"testing"

	"github.com/gary-lgy/aoc2019/testutil"
)

func TestVmReturnValue(t *testing.T) {
	tc := []testutil.VMTC{
		{[]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, nil, 3500, nil},
		{[]int64{1, 0, 0, 0, 99}, nil, 2, nil},
		{[]int64{2, 3, 0, 3, 99}, nil, 2, nil},
		{[]int64{2, 4, 4, 5, 99, 0}, nil, 2, nil},
		{[]int64{1, 1, 1, 4, 99, 5, 6, 0, 99}, nil, 30, nil},
	}
	c1 := testutil.ReadIntcodeFromFile(t, filepath.Join("input", "2.txt"))
	c2 := make([]int64, len(c1))
	copy(c2, c1)
	c1[1], c1[2] = 12, 2
	c2[1], c2[2] = 80, 18
	tc = append(tc,
		testutil.VMTC{Program: c1, Input: nil, ExpectedReturnValue: 3166704, ExpectedOutput: nil},
		testutil.VMTC{Program: c2, Input: nil, ExpectedReturnValue: 19690720, ExpectedOutput: nil})

	testutil.IntcodeVMTest(t, tc)
}
