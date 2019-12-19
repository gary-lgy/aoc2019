package intcode

import "fmt"

/* Possible modes:
0: position mode (represents a memory address)
1: immediate mode
2: relative mode (represents a memory address that to be be added the relative base)
*/
type parameter struct {
	mode  int
	value int64
}

func (param *parameter) getValue(vm *VM) (int64, error) {
	switch param.mode {
	case 0:
		return vm.readMemory(param.value), nil
	case 1:
		return param.value, nil
	case 2:
		return vm.readMemory(param.value + vm.relativeBase), nil
	default:
		return 0, fmt.Errorf("intcode VM error: unknown parameter mode for parameter value")
	}
}

func (param *parameter) getAddress(vm *VM) (int64, error) {
	switch param.mode {
	case 0:
		return param.value, nil
	case 2:
		return param.value + vm.relativeBase, nil
	default:
		return 0, fmt.Errorf("intcode VM error: unknown parameter mode for memory address")
	}
}
