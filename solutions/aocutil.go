package aoc2019

// Check panics if e is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// IntPair is a pair of int
type IntPair struct {
	X, Y int
}
