package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"sort"
	"strings"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solvers["10a"] = solve10a
	solvers["10b"] = solve10b
}

func readAsteroids(data string) []aocutil.IntPair {
	var asteroids []aocutil.IntPair
	for i, line := range strings.Split(strings.TrimSpace(data), "\n") {
		for j, ch := range line {
			if ch == '#' {
				asteroids = append(asteroids, aocutil.IntPair{X: j, Y: i})
			}
		}
	}
	return asteroids
}

func reduceFraction(frac aocutil.IntPair) aocutil.IntPair {
	if frac.X == 0 {
		return aocutil.IntPair{X: 0, Y: frac.Y / aocutil.AbsInt(frac.Y)}
	}
	if frac.Y == 0 {
		return aocutil.IntPair{X: frac.X / aocutil.AbsInt(frac.X), Y: 0}
	}
	gcd := aocutil.GreatestCommonDivisor(int64(aocutil.AbsInt(frac.X)), int64(aocutil.AbsInt(frac.Y)))
	return aocutil.IntPair{X: frac.X / int(gcd), Y: frac.Y / int(gcd)}
}

func bestMonitoringStation(asteroids []aocutil.IntPair) (aocutil.IntPair, int) {
	maxCount, station := math.MinInt32, aocutil.IntPair{}
	for _, current := range asteroids {
		count := 0
		visible := make(map[aocutil.IntPair]bool)
		for _, other := range asteroids {
			if current == other {
				continue
			}
			diff := reduceFraction(aocutil.IntPair{X: current.X - other.X, Y: current.Y - other.Y})
			if _, exists := visible[diff]; exists {
				continue
			}
			count++
			visible[diff] = true
		}
		if count > maxCount {
			maxCount = count
			station = current
		}
	}
	return station, maxCount
}

func solve10a(input io.Reader) (string, error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}
	asteroids := readAsteroids(string(data))
	_, ans := bestMonitoringStation(asteroids)
	return fmt.Sprint(ans), nil
}

func compareDirections(a, b aocutil.IntPair) bool {
	aHypo, bHypo := math.Sqrt(float64(a.X*a.X+a.Y*a.Y)), math.Sqrt(float64(b.X*b.X+b.Y*b.Y))
	aCosine, bCosine := -float64(a.Y)/aHypo, -float64(b.Y)/bHypo
	switch {
	case a.X >= 0 && b.X >= 0:
		return aCosine > bCosine
	case a.X < 0 && b.X < 0:
		return aCosine < bCosine
	case a.X >= 0 && b.X < 0:
		return true
	case a.X < 0 && b.X >= 0:
		return false
	default:
		panic("Bug in compareDirection! Not all the cases are considered.")
	}
}

func find200thVaporizedAsteroid(station aocutil.IntPair, asteroids []aocutil.IntPair) (aocutil.IntPair, error) {
	others := make(map[aocutil.IntPair][]aocutil.IntPair)
	for _, other := range asteroids {
		if station == other {
			continue
		}

		diff := aocutil.IntPair{X: other.X - station.X, Y: other.Y - station.Y}
		direction := reduceFraction(diff)
		others[direction] = append(others[direction], diff)
	}
	// Collect all directions and sort them. Also sort the targets in each direction
	var directions []aocutil.IntPair
	for dir, targets := range others {
		directions = append(directions, dir)
		sort.Slice(targets, func(i, j int) bool {
			a, b := targets[i], targets[j]
			return a.X*a.X+a.Y*a.Y < b.X*b.X+b.Y*b.Y
		})
	}
	sort.Slice(directions, func(i, j int) bool {
		return compareDirections(directions[i], directions[j])
	})
	i := 0
	for i < len(asteroids)-1 {
		for _, dir := range directions {
			if len(others[dir]) > 0 {
				i++
				diff := others[dir][0]
				if i == 200 {
					return aocutil.IntPair{X: station.X + diff.X, Y: station.Y + diff.Y}, nil
				}
				others[dir] = others[dir][1:]
			}
		}
	}
	return aocutil.IntPair{}, fmt.Errorf("less than 200 asteroids to vaporize")
}

func solve10b(input io.Reader) (string, error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}
	asteroids := readAsteroids(string(data))
	station, _ := bestMonitoringStation(asteroids)
	vaporizedAsteroid, err := find200thVaporizedAsteroid(station, asteroids)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(vaporizedAsteroid), nil
}
