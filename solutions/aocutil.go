package aoc2019

// check panics if e is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
