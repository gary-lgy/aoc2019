package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type adjList map[string][]string

func init() {
	solverMap["6a"] = solve6a
	solverMap["6b"] = solve6b
}

func readOrbits(input io.Reader) (orbits adjList) {
	scanner := bufio.NewScanner(input)
	orbits = make(adjList)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ")")
		from, to := line[0], line[1]
		orbits[from] = append(orbits[from], to)
	}
	return
}

func countOrbits(start string, orbits *adjList) (size, orbitCount int) {
	size, orbitCount = 1, 0
	for _, child := range (*orbits)[start] {
		s, c := countOrbits(child, orbits)
		size, orbitCount = size+s, orbitCount+c
	}
	orbitCount += size - 1
	return
}

func solve6a(input *os.File) {
	orbits := readOrbits(input)
	_, orbitCount := countOrbits("COM", &orbits)
	fmt.Println(orbitCount)
}

func getParents(input io.Reader) (parents map[string]string) {
	orbits := readOrbits(input)
	parents = make(map[string]string)
	for from, tos := range orbits {
		for _, to := range tos {
			parents[to] = from
		}
	}
	return
}

func backTrace(parents *map[string]string, start, end string) map[string]bool {
	path := make(map[string]bool)
	current := (*parents)[start]
	for current != end {
		path[current] = true
		current = (*parents)[current]
	}
	return path
}

func calcDist(parents *map[string]string) int {
	you, santa := backTrace(parents, "YOU", "COM"), backTrace(parents, "SAN", "COM")
	dist := 0
	for node := range you {
		if _, exists := santa[node]; !exists {
			dist++
		}
	}
	for node := range santa {
		if _, exists := you[node]; !exists {
			dist++
		}
	}
	return dist
}

func solve6b(input *os.File) {
	parents := getParents(input)
	fmt.Println(calcDist(&parents))
}
