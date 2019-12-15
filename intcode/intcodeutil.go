package intcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	. "github.com/gary-lgy/aoc2019/aocutil"
)

type parameter struct {
	mode, value int
}

func (param *parameter) getValue(memory []int) int {
	switch param.mode {
	case 0:
		return memory[param.value]
	case 1:
		return param.value
	default:
		panic("Unknown parameter mode")
	}
}

func numberOfParameters(opcode int) int {
	switch opcode {
	case 1, 2, 7, 8:
		return 3
	case 3, 4:
		return 1
	case 5, 6:
		return 2
	case 99:
		return 0
	default:
		panic("Unknown opcode " + strconv.Itoa(opcode))
	}
}

type instruction struct {
	opcode     int
	parameters []parameter
}

// execute inst and return whether the program should exit and the new PC, if applicable
// newPc will be negative when no jump should be performed
func (inst *instruction) execute(vm *Vm) (exit bool, newPc int) {
	switch inst.opcode {
	case 1:
		op1, op2, dest := inst.parameters[0].getValue(vm.memory), inst.parameters[1].getValue(vm.memory), inst.parameters[2].value
		vm.memory[dest] = op1 + op2
		return false, -1
	case 2:
		op1, op2, dest := inst.parameters[0].getValue(vm.memory), inst.parameters[1].getValue(vm.memory), inst.parameters[2].value
		vm.memory[dest] = op1 * op2
		return false, -1
	case 3:
		dest := inst.parameters[0].value
		vm.memory[dest] = vm.getInput()
		return false, -1
	case 4:
		out := inst.parameters[0].getValue(vm.memory)
		vm.pushOutput(out)
		return false, -1
	case 5:
		if inst.parameters[0].getValue(vm.memory) != 0 {
			return false, inst.parameters[1].getValue(vm.memory)
		} else {
			return false, -1
		}
	case 6:
		if inst.parameters[0].getValue(vm.memory) == 0 {
			return false, inst.parameters[1].getValue(vm.memory)
		} else {
			return false, -1
		}
	case 7:
		op1, op2, dest := inst.parameters[0].getValue(vm.memory), inst.parameters[1].getValue(vm.memory), inst.parameters[2].value
		if op1 < op2 {
			vm.memory[dest] = 1
		} else {
			vm.memory[dest] = 0
		}
		return false, -1
	case 8:
		op1, op2, dest := inst.parameters[0].getValue(vm.memory), inst.parameters[1].getValue(vm.memory), inst.parameters[2].value
		if op1 == op2 {
			vm.memory[dest] = 1
		} else {
			vm.memory[dest] = 0
		}
		return false, -1
	case 99:
		return true, -1
	default:
		panic("Unknown opcode " + strconv.Itoa(inst.opcode))
	}
}

func readOpcode(token int) (opcode int, modes []int) {
	opcode, token = token%100, token/100
	modes = make([]int, numberOfParameters(opcode))
	for i, _ := range modes {
		modes[i], token = token%10, token/10
	}
	return
}

type Vm struct {
	memory        []int
	input, output []int
	pc            int
}

func ReadIntCode(input *os.File) (intcode []int) {
	buf, err := ioutil.ReadAll(input)
	Check(err)
	tokens := strings.Split(strings.TrimSpace(string(buf)), ",")
	intcode = make([]int, len(tokens))
	for i, token := range tokens {
		number, err := strconv.ParseInt(token, 10, 32)
		Check(err)
		intcode[i] = int(number)
	}
	return
}

func NewVm(intcodes, input []int) Vm {
	memory := make([]int, len(intcodes))
	copy(memory, intcodes)
	return Vm{memory: memory, input: input, output: []int{}, pc: 0}
}

func (vm *Vm) getInput() int {
	i := vm.input[0]
	vm.input = vm.input[1:]
	return i
}

func (vm *Vm) pushOutput(output int) {
	fmt.Println("Output from vm:", output)
	vm.output = append(vm.output, output)
}

func (vm *Vm) GetOutput() []int {
	return vm.output
}

func (vm *Vm) SetMemory(index, value int) {
	vm.memory[index] = value
}

func (vm *Vm) Run() int {
	for vm.pc < len(vm.memory) {
		opcode, modes := readOpcode(vm.memory[vm.pc])
		vm.pc += 1
		var parameters []parameter
		for _, mode := range modes {
			parameters = append(parameters, parameter{mode, vm.memory[vm.pc]})
			vm.pc += 1
		}
		inst := instruction{opcode, parameters}
		exit, newPc := inst.execute(vm)
		if exit {
			break
		}
		if newPc > 0 {
			vm.pc = newPc
		}
	}

	return vm.memory[0]
}
