package main

import (
	"testing"

	. "github.com/gary-lgy/aoc2019/testutil"
)

func TestFuelPartA(t *testing.T) {
	testCases := []IntTC{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range testCases {
		actual := fuelPartA(c.Input)
		if c.Expected != actual {
			t.Errorf("fuelPartA(%d): expected %d, got %d", c.Input, c.Expected, actual)
		}
	}
}

func TestFuelPartB(t *testing.T) {
	testCases := []IntTC{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, c := range testCases {
		actual := fuelPartB(c.Input)
		if c.Expected != actual {
			t.Errorf("fuelPartA(%d): expected %d, got %d", c.Input, c.Expected, actual)
		}
	}
}
