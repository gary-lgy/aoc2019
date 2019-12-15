package testutil

type IntTc struct {
	Input, Expected int
}

type VmTc struct {
	Program, Input      []int
	ExpectedReturnValue int
	ExpectedOutput      []int
}
