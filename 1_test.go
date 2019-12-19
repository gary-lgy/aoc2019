package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type intTC struct {
	Input, Expected int
}

func TestFuelPartA(t *testing.T) {
	testCases := []intTC{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range testCases {
		assert.Equal(t, c.Expected, fuelPartA(c.Input))
	}
}

func TestFuelPartB(t *testing.T) {
	testCases := []intTC{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range testCases {
		assert.Equal(t, c.Expected, fuelPartB(c.Input))
	}
}
