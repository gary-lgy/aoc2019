package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
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

// RunSingleInstance constructs and runs a single VM with program and input, and return its output and exit code
func RunSingleInstance(program, input []int64, outputHandlers ...func(int64)) (output []int64, exitCode int64) {
	ic, oc := make(chan int64), make(chan int64)
	vm := NewVM(program, ic, oc)
	go vm.Run()
	for _, in := range input {
		ic <- in
	}
	for out := range oc {
		output = append(output, out)
		for _, fn := range outputHandlers {
			fn(out)
		}
	}
	return output, vm.ExitCode()
}

// RunWithWG runs vm and decrement wg count after it is done
func  RunWithWG(vm *VM, wg *sync.WaitGroup) {
	vm.Run()
	wg.Done()
}

