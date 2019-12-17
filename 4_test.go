package main

import "testing"

type passwordTc struct {
	Password int
	Ok       bool
}

func TestIsPossiblePasswordA(t *testing.T) {
	tc := []passwordTc{
		{111111, true},
		{223450, false},
		{123789, false},
	}
	for _, c := range tc {
		actual := isPossiblePartA(c.Password)
		if actual != c.Ok {
			t.Errorf("Expected isPossiblePartA(%d) to be %v, got %v", c.Password, c.Ok, actual)
		}
	}
}

func TestIsPossiblePasswordB(t *testing.T) {
	tc := []passwordTc{
		{112233, true},
		{123444, false},
		{111122, true},
	}
	for _, c := range tc {
		actual := isPossiblePartB(c.Password)
		if actual != c.Ok {
			t.Errorf("Expected isPossiblePartB(%d) to be %v, got %v", c.Password, c.Ok, actual)
		}
	}
}
