package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

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

// LogOutput logs output from intcode vm
func LogOutput(output int64) {
	fmt.Println("Output from VM:", output)
}

// RunSingleInstance constructs and runs a single VM with program and input, and returns its output, exit code, and error encountered, if any
func RunSingleInstance(program, input []int64) ([]int64, int64, error) {
	vm := NewVM(program)
	output, err := vm.Run(input)
	if err != nil {
		return nil, 0, err
	}
	return output, vm.ExitCode(), nil
}
