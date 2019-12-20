package intcode

import (
	"fmt"
)

// VM represents an intcode machine
// New VM instances should be constructed with the NewVM method instead of direct initialization of the VM struct
type VM struct {
	pc           int64
	relativeBase int64
	memory       map[int64]int64
	done         bool
	input        []int64
	output       []int64
}

// NewVM constructs a new Intcode VM
func NewVM(intcodes []int64) *VM {
	memory := make(map[int64]int64, len(intcodes))
	for i, c := range intcodes {
		memory[int64(i)] = c
	}
	return &VM{pc: 0, relativeBase: 0, memory: memory, done: false}
}

func (vm *VM) getInput() (int64, error) {
	if len(vm.input) == 0 {
		return 0, fmt.Errorf("not enough input")
	}
	in := vm.input[0]
	vm.input = vm.input[1:]
	return in, nil
}

func (vm *VM) pushOutput(output int64) {
	vm.output = append(vm.output, output)
}

func (vm *VM) readMemory(index int64) int64 {
	value, ok := vm.memory[index]
	if !ok {
		return 0
	}
	return value
}

// setMemory sets the memory of vm at index to value
func (vm *VM) setMemory(index, value int64) {
	vm.memory[index] = value
}

func (vm *VM) setPC(newPC int64) {
	vm.pc = newPC
}

func (vm *VM) adjustRelativeBase(change int64) {
	vm.relativeBase += change
}

// execute inst and return whether the memory should exit and the error encountered, if applicable
func (vm *VM) execute(inst *instruction) error {
	switch inst.opcode {
	case 1:
		return vm.executeArithmetic(inst,
			func(op1, op2 int64) int64 { return op1 + op2 })
	case 2:
		return vm.executeArithmetic(inst,
			func(op1, op2 int64) int64 { return op1 * op2 })
	case 3:
		return vm.executeInput(inst)
	case 4:
		return vm.executeOutput(inst)
	case 5:
		return vm.executeJumpIf(inst, func(value int64) bool { return value != 0 })
	case 6:
		return vm.executeJumpIf(inst, func(value int64) bool { return value == 0 })
	case 7:
		return vm.executeSetIf(inst, func(op1, op2 int64) bool { return op1 < op2 })
	case 8:
		return vm.executeSetIf(inst, func(op1, op2 int64) bool { return op1 == op2 })
	case 9:
		return vm.executeChangeRelativeBase(inst)
	case 99:
		vm.done = true
		return nil
	default:
		return fmt.Errorf("unknown opcode %d", inst.opcode)
	}
}

// Run vm with inputs until more inputs are needed or the program halts
func (vm *VM) Run(input []int64) ([]int64, error) {
	vm.input = make([]int64, len(input))
	copy(vm.input, input)
	vm.output = nil

	for !vm.done {
		opcode, modes := readOpcode(vm.memory[vm.pc])
		if opcode == 3 && len(vm.input) == 0 {
			break // no more input
		}
		vm.pc++
		var parameters []parameter
		for _, mode := range modes {
			parameters = append(parameters, parameter{mode, vm.memory[vm.pc]})
			vm.pc++
		}
		inst := instruction{opcode, parameters}
		err := vm.execute(&inst)
		if err != nil {
			return nil, fmt.Errorf("intcode VM internal error: %v", err)
		}
	}
	return vm.output, nil
}

// ExitCode returns the exit code of the VM
func (vm *VM) ExitCode() int64 {
	return vm.memory[0]
}

// Stopped checks if vm has stopped running
func (vm *VM) Stopped() bool {
	return vm.done
}
