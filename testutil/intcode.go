package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gary-lgy/aoc2019/intcode"
)

// IntcodeVMTest tests VM with test cases tc
func IntcodeVMTest(t *testing.T, tc []VMTC) {
	for _, c := range tc {
		ic, oc := make(chan int), make(chan int)
		vm := intcode.NewVM(c.Program, ic, oc)
		go vm.Run()
		for _, i := range c.Input {
			ic <- i
		}
		actualOutput := make([]int, 0)
		for o := range oc {
			actualOutput = append(actualOutput, o)
		}
		exitCode := vm.ExitCode()
		assert.Equalf(t, c.ExpectedReturnValue, exitCode,
			"Running %v with input %v, expected %d, got %d", c.Program, c.Input, c.ExpectedReturnValue, exitCode)
		assert.Equalf(t, c.ExpectedOutput, actualOutput,
			"Running %v with input %v, expected output %v, got %v", c.Program, c.Input, c.ExpectedOutput, actualOutput)
	}
}
