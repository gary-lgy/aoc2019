package intcode

func (inst *instruction) executeArithmetic(vm *VM, fn func(int64, int64) int64) (bool, int64, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	dest, err := inst.parameters[2].getAddress(vm)
	if err != nil {
		return false, -1, err
	}
	vm.SetMemory(dest, fn(op1, op2))
	return false, -1, nil
}

func (inst *instruction) executeInput(vm *VM) (bool, int64, error) {
	dest, err := inst.parameters[0].getAddress(vm)
	if err != nil {
		return false, -1, err
	}
	vm.SetMemory(dest, vm.getInput())
	return false, -1, nil
}

func (inst *instruction) executeOutput(vm *VM) (bool, int64, error) {
	out, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	vm.pushOutput(out)
	return false, -1, nil
}

func (inst *instruction) executeJumpIf(vm *VM, predicate func(int64) bool) (bool, int64, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	if predicate(op1) {
		return false, op2, nil
	}
	return false, -1, nil
}

func (inst *instruction) executeSetIf(vm *VM, predicate func(int64, int64) bool) (bool, int64, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	dest, err := inst.parameters[2].getAddress(vm)
	if err != nil {
		return false, -1, err
	}
	if predicate(op1, op2) {
		vm.SetMemory(dest, 1)
	} else {
		vm.SetMemory(dest, 0)
	}
	return false, -1, nil
}

func (inst *instruction) executeChangeRelativeBase(vm *VM) (bool, int64, error) {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return false, -1, err
	}
	vm.relativeBase += op1
	return false, -1, nil
}
