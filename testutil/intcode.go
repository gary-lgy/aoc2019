package testutil

import (
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
)

// IntcodeVMTest tests VM with testcases tc
func IntcodeVMTest(t *testing.T, tc []VMTC) {
	for _, c := range tc {
		vm := intcode.NewVM(c.Program, c.Input)
		ret := vm.Run()
		output := vm.GetOutput()
		if ret != c.ExpectedReturnValue {
			t.Errorf("Running %v with input %v, expected %d, got %d", c.Program, c.Input, c.ExpectedReturnValue, ret)
		}
		if !aocutil.IntSliceEqual(c.ExpectedOutput, output) {
			t.Errorf("Running %v with input %v, expected output %v, got %v", c.Program, c.Input, c.ExpectedOutput, output)
		}
	}
}
