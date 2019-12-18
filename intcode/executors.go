package intcode

func (inst *instruction) executeArithmetic(vm *VM, fn func(int, int) int) (bool, int, error) {
	op1, err := inst.parameters[0].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	op2, err := inst.parameters[1].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	dest := inst.parameters[2].value
	vm.memory[dest] = fn(op1, op2)
	return false, -1, nil
}

func (inst *instruction) executeInput(vm *VM) (bool, int, error) {
	dest := inst.parameters[0].value
	vm.memory[dest] = vm.getInput()
	return false, -1, nil
}

func (inst *instruction) executeOutput(vm *VM) (bool, int, error) {
	out, err := inst.parameters[0].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	vm.pushOutput(out)
	return false, -1, nil
}

func (inst *instruction) executeJumpIf(vm *VM, predicate func(int) bool) (bool, int, error) {
	op1, err := inst.parameters[0].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	op2, err := inst.parameters[1].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	if predicate(op1) {
		return false, op2, nil
	}
	return false, -1, nil
}

func (inst *instruction) executeSetIf(vm *VM, predicate func(int, int) bool) (bool, int, error) {
	op1, err := inst.parameters[0].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	op2, err := inst.parameters[1].getValue(vm.memory)
	if err != nil {
		return false, -1, err
	}
	dest := inst.parameters[2].value
	if predicate(op1, op2) {
		vm.memory[dest] = 1
	} else {
		vm.memory[dest] = 0
	}
	return false, -1, nil
}
