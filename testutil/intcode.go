package testutil

import (
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
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
		var actualOutput []int
		for o := range oc {
			actualOutput = append(actualOutput, o)
		}
		if vm.ExitCode() != c.ExpectedReturnValue {
			t.Errorf("Running %v with input %v, expected %d, got %d", c.Program, c.Input, c.ExpectedReturnValue, vm.ExitCode())
		}
		if !aocutil.IntSliceEqual(c.ExpectedOutput, actualOutput) {
			t.Errorf("Running %v with input %v, expected output %v, got %v", c.Program, c.Input, c.ExpectedOutput, actualOutput)
		}
	}
}
