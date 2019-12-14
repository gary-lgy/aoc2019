package aoc2019

import (
	"math"
)

// AbsInt returns the absoute value of x
func AbsInt(x int) int {
	if (x < 0) {
		return -x
	} else {
		return x
	}
}

// MaxInt returns the maximum among the arguments
func MaxInt(ints ...int) int {
	max := math.MinInt32
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

// MinInt returns the minimum among the arguments
func MinInt(ints ...int) int {
	min := math.MaxInt32
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}
