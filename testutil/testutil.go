package testutil

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gary-lgy/aoc2019/intcode"
)

// VMTC is a test case for Intcode VM
type VMTC struct {
	Program, Input      []int64
	ExpectedReturnValue int64
	ExpectedOutput      []int64
}

// ReadIntcodeFromFile reads and returns an Intcode program from the file given by filepath.
// It fails the current test if reading from the file fails.
func ReadIntcodeFromFile(t *testing.T, filepath string) []int64 {
	input, err := os.Open(filepath)
	require.NoError(t, err)
	defer input.Close()
	program, err := intcode.ReadIntCode(input)
	require.NoError(t, err)
	return program
}

// ReadStringFromFile reads and returns the content of the file given by filepath.
// It fails the current test if reading from the file fails.
func ReadStringFromFile(t *testing.T, filepath string) string {
	input, err := os.Open(filepath)
	require.NoError(t, err)
	defer input.Close()
	data, err := ioutil.ReadAll(input)
	require.NoError(t, err)
	return string(data)
}
