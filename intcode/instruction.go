package intcode

import "fmt"

/*
Permitted instructions:
1. Addition             (memory[op3] = op1 + op2)
2. Multiplication       (memory[op3] = op1 + op2)
3. Input                (memory[op1] = input)
4. Output               (memory[op1])
5. Jump-If-True         (jump if op1 != 0)
6. Jump-if-False        (jump if op1 == 0)
7. Set-Less-Than        (memory[op3] = 1 if op1 < op2 else 0)
8. Set-Equal            (memory[op3] = 1 if op1 == op2 else 0)
9. Change-Relative-Base (RB += op1)
*/
type instruction struct {
	opcode     int
	parameters []parameter
}

func numberOfParameters(opcode int) (int, error) {
	switch opcode {
	case 1, 2, 7, 8:
		return 3, nil
	case 3, 4, 9:
		return 1, nil
	case 5, 6:
		return 2, nil
	case 99:
		return 0, nil
	default:
		return 0, fmt.Errorf("unknown opcode %d", opcode)
	}
}

func readOpcode(token int64) (opcode int, modes []int) {
	opcode, token = int(token%100), token/100
	num, err := numberOfParameters(opcode)
	if err != nil {
		return 0, nil
	}
	modes = make([]int, num)
	for i := range modes {
		modes[i], token = int(token%10), token/10
	}
	return
}
