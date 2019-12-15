package testutil

import (
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
)

func IntcodeVmTest(t *testing.T, tc []VmTc) {
	for _, c := range tc {
		vm := intcode.NewVm(c.Program, c.Input)
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
