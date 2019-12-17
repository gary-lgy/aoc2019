package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
	"github.com/gary-lgy/aoc2019/intcode"
	. "github.com/gary-lgy/aoc2019/testutil"
)

func TestVm5a(t *testing.T) {
	input, err := os.Open(filepath.Join("input", "5"))
	defer input.Close()
	aocutil.Check(err)
	tc := VMTC{intcode.ReadIntCode(input), []int{1}, 3, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 4601506}}
	IntcodeVMTest(t, []VMTC{tc})
}
