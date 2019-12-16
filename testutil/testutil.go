package testutil

// IntTC is a test case that expectes an int output given an int input
type IntTC struct {
	Input, Expected int
}

// VMTC is a testcase for Intcode VM
type VMTC struct {
	Program, Input      []int
	ExpectedReturnValue int
	ExpectedOutput      []int
}
