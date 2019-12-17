package testutil

// IntTC is a test case that expects an int output given an int input
type IntTC struct {
	Input, Expected int
}

// VMTC is a test case for Intcode VM
type VMTC struct {
	Program, Input      []int
	ExpectedReturnValue int
	ExpectedOutput      []int
}

// StringIntTC is a test case that takes in a string and expects an int
type StringIntTC struct {
	Input    string
	Expected int
}
