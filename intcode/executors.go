package intcode

func (vm *VM) executeArithmetic(inst *instruction, fn func(int64, int64) int64) (bool, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return false, err
	}
	dest, err := inst.parameters[2].getAddress(vm)
	if err != nil {
		return false, err
	}
	vm.setMemory(dest, fn(op1, op2))
	return false, nil
}

func (vm *VM) executeInput(inst *instruction) (bool, error) {
	dest, err := inst.parameters[0].getAddress(vm)
	if err != nil {
		return false, err
	}
	vm.setMemory(dest, vm.getInput())
	return false, nil
}

func (vm *VM) executeOutput(inst *instruction) (bool, error) {
	out, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, err
	}
	vm.pushOutput(out)
	return false, nil
}

func (vm *VM) executeJumpIf(inst *instruction, predicate func(int64) bool) (bool, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return false, err
	}
	if predicate(op1) {
		vm.setPC(op2)
	}
	return false, nil
}

func (vm *VM) executeSetIf(inst *instruction, predicate func(int64, int64) bool) (bool, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return false, err
	}
	dest, err := inst.parameters[2].getAddress(vm)
	if err != nil {
		return false, err
	}
	if predicate(op1, op2) {
		vm.setMemory(dest, 1)
	} else {
		vm.setMemory(dest, 0)
	}
	return false, nil
}

func (vm *VM) executeChangeRelativeBase(inst *instruction) (bool, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, err
	}
	vm.adjustRelativeBase(op1)
	return false, nil
}
