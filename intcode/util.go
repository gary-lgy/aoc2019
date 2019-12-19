package intcode

import "fmt"

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
