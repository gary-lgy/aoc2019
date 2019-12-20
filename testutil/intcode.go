package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gary-lgy/aoc2019/intcode"
)

// IntcodeVMTest tests VM with test cases tc
func IntcodeVMTest(t *testing.T, tc []VMTC) {
	for _, c := range tc {
		actualOutput, exitCode, err := intcode.RunSingleInstance(c.Program, c.Input)
		require.NoError(t, err)
		assert.Equalf(t, c.ExpectedReturnValue, exitCode,
			"Running %v with input %v, expected %d, got %d", c.Program, c.Input, c.ExpectedReturnValue, exitCode)
		assert.Equalf(t, c.ExpectedOutput, actualOutput,
			"Running %v with input %v, expected output %v, got %v", c.Program, c.Input, c.ExpectedOutput, actualOutput)
	}
}
