package intcode

func (vm *VM) executeArithmetic(inst *instruction, fn func(int64, int64) int64) error {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return err
	}
	dest, err := inst.parameters[2].getAddress(vm)
	if err != nil {
		return err
	}
	vm.setMemory(dest, fn(op1, op2))
	return nil
}

func (vm *VM) executeInput(inst *instruction) error {
	dest, err := inst.parameters[0].getAddress(vm)
	if err != nil {
		return err
	}
	vm.setMemory(dest, vm.getInput())
	return nil
}

func (vm *VM) executeOutput(inst *instruction) error {
	out, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return err
	}
	vm.pushOutput(out)
	return nil
}

func (vm *VM) executeJumpIf(inst *instruction, predicate func(int64) bool) error {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return err
	}
	if predicate(op1) {
		vm.setPC(op2)
	}
	return nil
}

func (vm *VM) executeSetIf(inst *instruction, predicate func(int64, int64) bool) error {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return err
	}
	op2, err := inst.parameters[1].getValue(vm)
	if err != nil {
		return err
	}
	dest, err := inst.parameters[2].getAddress(vm)
	if err != nil {
		return err
	}
	if predicate(op1, op2) {
		vm.setMemory(dest, 1)
	} else {
		vm.setMemory(dest, 0)
	}
	return nil
}

func (vm *VM) executeChangeRelativeBase(inst *instruction) error {
	op1, err := inst.parameters[0].getValue(vm)
	if err != nil {
		return err
	}
	vm.adjustRelativeBase(op1)
	return nil
}
