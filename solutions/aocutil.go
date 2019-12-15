package aoc2019

// Check panics if e is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Ensure panics with reason if b is false
func Ensure(b bool)  {
	if !b {
		panic("Ensure failed")
	}
}

// IntPair is a pair of int
type IntPair struct {
	X, Y int
}
