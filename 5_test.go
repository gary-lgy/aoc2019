package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gary-lgy/aoc2019/intcode"
	. "github.com/gary-lgy/aoc2019/testutil"
)

func TestDay5(t *testing.T) {
	input, err := os.Open(filepath.Join("input", "5.txt"))
	require.NoError(t, err)
	defer input.Close()
	program, err := intcode.ReadIntCode(input)
	require.NoError(t, err)
	tc := []VMTC{
		{Program: program, Input: []int{1}, ExpectedReturnValue: 3, ExpectedOutput: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 4601506}},
		{Program: program, Input: []int{5}, ExpectedReturnValue: 314, ExpectedOutput: []int{5525561}},
	}
	IntcodeVMTest(t, tc)
}
