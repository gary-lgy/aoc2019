package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func test6a(t *testing.T, orbits *adjList, expected int) {
	_, actual := countOrbits("COM", orbits)
	if actual != expected {
		t.Errorf("Expected %d orbits, got %d", expected, actual)
	}
}

func Test6a(t *testing.T) {
	const example = "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
	orbits := readOrbits(strings.NewReader(example))
	test6a(t, &orbits, 42)

	input, err := os.Open(filepath.Join("input", "6"))
	aocutil.Check(err)
	defer input.Close()
	orbits = readOrbits(input)
	test6a(t, &orbits, 147807)
}

func test6b(t *testing.T, input io.Reader, expected int) {
	parents := getParents(input)
	actual := calcDist(&parents)
	if actual != expected {
		t.Errorf("Expected %d steps, got %d", expected, actual)
	}
}

func Test6b(t *testing.T) {
	const example = "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"

	test6b(t, strings.NewReader(example), 4)

	input, err := os.Open(filepath.Join("input", "6"))
	aocutil.Check(err)
	defer input.Close()
	test6b(t, input, 229)
}
