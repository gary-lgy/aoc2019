package intcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	. "github.com/gary-lgy/aoc2019/aocutil"
)

type Parameter struct {
	mode, value int
}

func (param *Parameter) Value(memory []int) int {
	if param.mode == 0 {
		return memory[param.value]
	} else if param.mode == 1 {
		return param.value
	} else {
		panic("Unknown parameter mode")
	}
}

type Instruction struct {
	opcode     int
	parameters []Parameter
}

// Execute inst and return whether the program should exit and the new PC, if applicable
// newPc will be negative when no jump should be performed
func (inst *Instruction) Execute(memory []int) (exit bool, newPc int) {
	switch inst.opcode {
	case 1:
		op1, op2, dest := inst.parameters[0].Value(memory), inst.parameters[1].Value(memory), inst.parameters[2].value
		memory[dest] = op1 + op2
		return false, -1
	case 2:
		op1, op2, dest := inst.parameters[0].Value(memory), inst.parameters[1].Value(memory), inst.parameters[2].value
		memory[dest] = op1 * op2
		return false, -1
	case 3:
		dest := inst.parameters[0].value
		var input int
		fmt.Scan(&input)
		memory[dest] = input
		return false, -1
	case 4:
		fmt.Println(inst.parameters[0].Value(memory))
		return false, -1
	case 5:
		if inst.parameters[0].Value(memory) != 0 {
			return false, inst.parameters[1].Value(memory)
		} else {
			return false, -1
		}
	case 6:
		if inst.parameters[0].Value(memory) == 0 {
			return false, inst.parameters[1].Value(memory)
		} else {
			return false, -1
		}
	case 7:
		op1, op2, dest := inst.parameters[0].Value(memory), inst.parameters[1].Value(memory), inst.parameters[2].value
		if op1 < op2 {
			memory[dest] = 1
		} else {
			memory[dest] = 0
		}
		return false, -1
	case 8:
		op1, op2, dest := inst.parameters[0].Value(memory), inst.parameters[1].Value(memory), inst.parameters[2].value
		if op1 == op2 {
			memory[dest] = 1
		} else {
			memory[dest] = 0
		}
		return false, -1
	case 99:
		return true, -1
	default:
		panic("Unknown opcode " + strconv.Itoa(inst.opcode))
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

func readOpcode(token int) (opcode int, modes []int) {
	opcode, token = token%100, token/100
	modes = make([]int, numberOfParameters(opcode))
	for i, _ := range modes {
		modes[i], token = token%10, token/10
	}
	return
}

func ReadProgram(input *os.File) (program []int) {
	buf, err := ioutil.ReadAll(input)
	Check(err)
	tokens := strings.Split(strings.TrimSpace(string(buf)), ",")
	for _, token := range tokens {
		number, err := strconv.ParseInt(token, 10, 32)
		Check(err)
		program = append(program, int(number))
	}
	return
}

func RunProgram(program []int) int {
	memory := make([]int, len(program))
	copy(memory, program)
	pc := 0
	for pc < len(memory) {
		opcode, modes := readOpcode(memory[pc])
		pc++
		var parameters []Parameter
		for _, mode := range modes {
			parameters = append(parameters, Parameter{mode, memory[pc]})
			pc++
		}
		instruction := Instruction{opcode, parameters}
		exit, newPc := instruction.Execute(memory)
		if exit {
			break
		}
		if newPc > 0 {
			pc = newPc
		}
	}

	return memory[0]
}

// TODO: virtual IO and tests
