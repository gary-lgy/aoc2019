package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

// TODO: Put in own file
/* Possible modes:
0: position mode (represents a memory address)
1: immediate mode
2: relative mode (represents a memory address that to be be added the relative base)
*/
type parameter struct {
	mode  int
	value int64
}

func numberOfParameters(opcode int) (int, error) {
	switch opcode {
	case 1, 2, 7, 8:
		return 3, nil
	case 3, 4, 9:
		return 1, nil
	case 5, 6:
		return 2, nil
	case 99:
		return 0, nil
	default:
		return 0, fmt.Errorf("unknown opcode %d", opcode)
	}
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

// TODO: Put in own file
/*
Permitted instructions:
1. Addition             (memory[op3] = op1 + op2)
2. Multiplication       (memory[op3] = op1 + op2)
3. Input                (memory[op1] = input)
4. Output               (memory[op1])
5. Jump-If-True         (jump if op1 != 0)
6. Jump-if-False        (jump if op1 == 0)
7. Set-Less-Than        (memory[op3] = 1 if op1 < op2 else 0)
8. Set-Equal            (memory[op3] = 1 if op1 == op2 else 0)
9. Change-Relative-Base (RB += op1)
*/
type instruction struct {
	opcode     int
	parameters []parameter
}

// TODO: refactor to be a method of VM
// execute inst and return whether the memory should exit and the new PC, if applicable
// newPc will be negative when no jump should be performed
func (inst *instruction) execute(vm *VM) (bool, int64, error) {
	switch inst.opcode {
	case 1:
		return inst.executeArithmetic(vm,
			func(op1, op2 int64) int64 { return op1 + op2 })
	case 2:
		return inst.executeArithmetic(vm,
			func(op1, op2 int64) int64 { return op1 * op2 })
	case 3:
		return inst.executeInput(vm)
	case 4:
		return inst.executeOutput(vm)
	case 5:
		return inst.executeJumpIf(vm, func(value int64) bool { return value != 0 })
	case 6:
		return inst.executeJumpIf(vm, func(value int64) bool { return value == 0 })
	case 7:
		return inst.executeSetIf(vm, func(op1, op2 int64) bool { return op1 < op2 })
	case 8:
		return inst.executeSetIf(vm, func(op1, op2 int64) bool { return op1 == op2 })
	case 9:
		return inst.executeChangeRelativeBase(vm)
	case 99:
		close(vm.output)
		return true, -1, nil
	default:
		return false, 0, fmt.Errorf("unknown opcode %d", inst.opcode)
	}
}

func readOpcode(token int64) (opcode int, modes []int) {
	opcode, token = int(token%100), token/100
	num, err := numberOfParameters(opcode)
	if err != nil {
		return 0, nil
	}
	modes = make([]int, num)
	for i := range modes {
		modes[i], token = int(token%10), token/10
	}
	return
}

// VM represents an intcode machine
// New VM instances should be constructed with the NewVM method instead of direct initialization of the VM struct
type VM struct {
	pc           int64
	relativeBase int64
	memory       map[int64]int64
	input        <-chan int64
	output       chan<- int64
	exitCode     int64
}

// TODO: Put in utils

// ReadIntCode takes in an io.Reader and returns its content as an intcode memory
func ReadIntCode(input io.Reader) ([]int64, error) {
	buf, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	tokens := strings.Split(strings.TrimSpace(string(buf)), ",")
	intcode := make([]int64, len(tokens))
	for i, token := range tokens {
		number, err := strconv.ParseInt(token, 10, 64)
		if err != nil {
			return nil, err
		}
		intcode[i] = number
	}
	return intcode, nil
}

// NewVM constructs a new Intcode VM
func NewVM(intcodes []int64, input <-chan int64, output chan<- int64) *VM {
	memory := make(map[int64]int64, len(intcodes))
	for i, c := range intcodes {
		memory[int64(i)] = c
	}
	return &VM{memory: memory, input: input, output: output, pc: 0, relativeBase: 0, exitCode: 0}
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

// SetMemory sets the memory of vm at index to value
func (vm *VM) SetMemory(index, value int64) {
	vm.memory[index] = value
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
		exit, newPc, err := inst.execute(vm)
		if err != nil {
			panic(fmt.Errorf("intcode VM internal error: %v", err))
		}
		if exit {
			break
		}
		// TODO: set pc in executor
		if newPc >= 0 {
			vm.pc = newPc
		}
	}
	vm.exitCode = vm.memory[0]
}

// RunWithWG runs vm and decrement wg count after it is done
func (vm *VM) RunWithWG(wg *sync.WaitGroup) {
	vm.Run()
	wg.Done()
}

// ExitCode returns the exit code of the VM
func (vm *VM) ExitCode() int64 {
	return vm.exitCode
}
