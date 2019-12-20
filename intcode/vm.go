package intcode

import (
	"fmt"
)

// TODO: Refactor to not use go routines?

// VM represents an intcode machine
// New VM instances should be constructed with the NewVM method instead of direct initialization of the VM struct
type VM struct {
	pc           int64
	relativeBase int64
	memory       map[int64]int64
	input        <-chan int64
	output       chan<- int64
	exitCode     int64
	done         bool
}

// NewVM constructs a new Intcode VM
func NewVM(intcodes []int64, input <-chan int64, output chan<- int64) *VM {
	memory := make(map[int64]int64, len(intcodes))
	for i, c := range intcodes {
		memory[int64(i)] = c
	}
	return &VM{memory: memory, input: input, output: output, pc: 0, relativeBase: 0, exitCode: 0, done: false}
}

func (vm *VM) getInput() int64 {
	return <-vm.input
}

func (vm *VM) pushOutput(output int64) {
	vm.output <- output
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
		vm.exitCode = vm.memory[0]
		close(vm.output)
		return nil
	default:
		return fmt.Errorf("unknown opcode %d", inst.opcode)
	}
}

// Run vm
func (vm *VM) Run() {
	for {
		opcode, modes := readOpcode(vm.memory[vm.pc])
		vm.pc++
		var parameters []parameter
		for _, mode := range modes {
			parameters = append(parameters, parameter{mode, vm.memory[vm.pc]})
			vm.pc++
		}
		inst := instruction{opcode, parameters}
		err := vm.execute(&inst)
		if err != nil {
			panic(fmt.Errorf("intcode VM internal error: %v", err))
		}
		if vm.done {
			break
		}
	}
}

// ExitCode returns the exit code of the VM
func (vm *VM) ExitCode() int64 {
	return vm.exitCode
}

// Stopped checks if vm has stopped running
func (vm *VM) Stopped() bool {
	return vm.done
}
