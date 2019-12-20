package aocutil

import (
	"math"
)

// IntPair is a pair of int
type IntPair struct {
	X, Y int
}

// AbsInt returns the absolute value of x
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MaxInt returns the maximum among the arguments
func MaxInt(integers ...int) int {
	max := math.MinInt32
	for _, i := range integers {
		if i > max {
			max = i
		}
	}
	return max
}

// MinInt returns the minimum among the arguments
func MinInt(integers ...int) int {
	min := math.MaxInt32
	for _, i := range integers {
		if i < min {
			min = i
		}
	}
	return min
}

// GreatestCommonDivisor computes the gcd of a and b
func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}
