package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
)

type wireTc struct {
	Input    string
	Expected int
}

func Test3a(t *testing.T) {
	tc := []wireTc{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`, 135},
	}
	input, err := os.Open(filepath.Join("input", "3"))
	aocutil.Check(err)
	defer input.Close()
	data, err := ioutil.ReadAll(input)
	aocutil.Check(err)
	tc = append(tc, wireTc{string(data), 1084})

	for _, c := range tc {
		actual := shortestManhattanDistance(parseWires(c.Input))
		if actual != c.Expected {
			t.Errorf("shortestManhattanDistance(%q) = %d, expected %d", c.Input, actual, c.Expected)
		}
	}
}

func Test3b(t *testing.T) {
	tc := []wireTc{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 610},
		{`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`, 410},
	}
	input, err := os.Open(filepath.Join("input", "3"))
	aocutil.Check(err)
	defer input.Close()
	data, err := ioutil.ReadAll(input)
	aocutil.Check(err)
	tc = append(tc, wireTc{string(data), 9240})

	for _, c := range tc {
		actual := shortestDelay(parseWires(c.Input))
		if actual != c.Expected {
			t.Errorf("shortestDelay(%q) = %d, expected %d", c.Input, actual, c.Expected)
		}
	}
}
