package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

type parameter struct {
	mode, value int
}

func (param *parameter) getValue(memory []int) (int, error) {
	switch param.mode {
	case 0:
		return memory[param.value], nil
	case 1:
		return param.value, nil
	default:
		return 0, fmt.Errorf("intcode VM error: unknown parameter mode")
	}
}

func numberOfParameters(opcode int) (int, error) {
	switch opcode {
	case 1, 2, 7, 8:
		return 3, nil
	case 3, 4:
		return 1, nil
	case 5, 6:
		return 2, nil
	case 99:
		return 0, nil
	default:
		return 0, fmt.Errorf("unknown opcode %d", opcode)
	}
}

type instruction struct {
	opcode     int
	parameters []parameter
}

// execute inst and return whether the program should exit and the new PC, if applicable
// newPc will be negative when no jump should be performed
func (inst *instruction) execute(vm *VM) (bool, int, error) {
	switch inst.opcode {
	case 1:
		return inst.executeArithmetic(vm,
			func(op1, op2 int) int { return op1 + op2 })
	case 2:
		return inst.executeArithmetic(vm,
			func(op1, op2 int) int { return op1 * op2 })
	case 3:
		return inst.executeInput(vm)
	case 4:
		return inst.executeOutput(vm)
	case 5:
		return inst.executeJumpIf(vm, func(value int) bool { return value != 0 })
	case 6:
		return inst.executeJumpIf(vm, func(value int) bool { return value == 0 })
	case 7:
		return inst.executeSetIf(vm, func(op1, op2 int) bool { return op1 < op2 })
	case 8:
		return inst.executeSetIf(vm, func(op1, op2 int) bool { return op1 == op2 })
	case 99:
		close(vm.output)
		return true, -1, nil
	default:
		return false, 0, fmt.Errorf("unknown opcode %d", inst.opcode)
	}
}

func readOpcode(token int) (opcode int, modes []int) {
	opcode, token = token%100, token/100
	num, err := numberOfParameters(opcode)
	if err != nil {
		return 0, nil
	}
	modes = make([]int, num)
	for i := range modes {
		modes[i], token = token%10, token/10
	}
	return
}

// VM represents an intcode machine
type VM struct {
	pc       int
	memory   []int
	input    <-chan int
	output   chan<- int
	exitCode int
}

// ReadIntCode takes in an io.Reader and returns its content as an intcode program
func ReadIntCode(input io.Reader) ([]int, error) {
	buf, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	tokens := strings.Split(strings.TrimSpace(string(buf)), ",")
	intcode := make([]int, len(tokens))
	for i, token := range tokens {
		number, err := strconv.ParseInt(token, 10, 32)
		if err != nil {
			return nil, err
		}
		intcode[i] = int(number)
	}
	return intcode, nil
}

// NewVM constructs a new Intcode VM
func NewVM(intcodes []int, input <-chan int, output chan<- int) VM {
	memory := make([]int, len(intcodes))
	copy(memory, intcodes)
	return VM{memory: memory, input: input, output: output, pc: 0, exitCode: 0}
}

func (vm *VM) getInput() int {
	return <-vm.input
}

func (vm *VM) pushOutput(output int) {
	vm.output <- output
}

// SetMemory sets the memory of vm at index to value
func (vm *VM) SetMemory(index, value int) {
	vm.memory[index] = value
}

// Run vm
func (vm *VM) Run() {
	for vm.pc < len(vm.memory) {
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
		if newPc > 0 {
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
func (vm *VM) ExitCode() int {
	return vm.exitCode
}
