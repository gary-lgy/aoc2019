package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solverMap["3a"] = solve3a
	solverMap["3b"] = solve3b
}

type point aocutil.IntPair

func (p *point) manhattanDistance(other *point) int {
	return aocutil.AbsInt(other.X-p.X) + aocutil.AbsInt(other.Y-p.Y)
}

type segment struct {
	start, end             point
	xMin, xMax, yMin, yMax int
}

func (s *segment) length() int {
	return s.start.manhattanDistance(&s.end)
}

func newSegment(start, end point) segment {
	s := segment{start: start, end: end}
	s.xMin = aocutil.MinInt(start.X, end.X)
	s.xMax = aocutil.MaxInt(start.X, end.X)
	s.yMin = aocutil.MinInt(start.Y, end.Y)
	s.yMax = aocutil.MaxInt(start.Y, end.Y)
	return s
}

func getSegments(wire string) []segment {
	var segments []segment
	descriptors := strings.Split(wire, ",")
	start := point{0, 0}
	var end point
	for _, desc := range descriptors {
		direction := desc[0]
		steps, err := strconv.ParseInt(desc[1:], 10, 32)
		aocutil.Check(err)
		switch direction {
		case 'U':
			end = point{start.X, start.Y + int(steps)}
		case 'D':
			end = point{start.X, start.Y - int(steps)}
		case 'L':
			end = point{start.X - int(steps), start.Y}
		case 'R':
			end = point{start.X + int(steps), start.Y}
		default:
			panic("Unknown direction " + string(direction))
		}
		segments = append(segments, newSegment(start, end))
		start = end
	}
	return segments
}

func parseWires(input string) ([]segment, []segment) {
	wires := strings.Split(strings.TrimSpace(input), "\n")
	return getSegments(wires[0]), getSegments(wires[1])
}

func intersection(s1, s2 segment) (intersection point, exists bool) {
	switch {
	case s1.xMin <= s2.xMin && s2.xMax <= s1.xMax && s2.yMin <= s1.yMin && s1.yMax <= s2.yMax:
		intersection = point{s2.xMin, s1.yMin}
		exists = true
	case s2.xMin <= s1.xMin && s1.xMax <= s2.xMax && s1.yMin <= s2.yMin && s2.yMax <= s1.yMax:
		intersection = point{s1.xMin, s2.yMin}
		exists = true
	default:
		intersection = point{0, 0} // Dummy point
		exists = false
	}
	return
}

func shortestManhattanDistance(w1, w2 []segment) int {
	origin := point{0, 0}
	min := math.MaxInt32
	for _, segment1 := range w1 {
		for _, segment2 := range w2 {
			if intersection, exists := intersection(segment1, segment2); exists && intersection != (point{0, 0}) {
				min = aocutil.MinInt(min, origin.manhattanDistance(&intersection))
			}
		}
	}
	return min
}

func solve3a(input *os.File) {
	buf, err := ioutil.ReadAll(input)
	aocutil.Check(err)
	w1, w2 := parseWires(string(buf))
	fmt.Println(shortestManhattanDistance(w1, w2))
}

func shortestDelay(w1, w2 []segment) int {
	min := math.MaxInt32
	d1 := 0
	for _, segment1 := range w1 {
		d2 := 0
		for _, segment2 := range w2 {
			if intersection, exists := intersection(segment1, segment2); exists && intersection != (point{0, 0}) {
				min = aocutil.MinInt(min, d1+d2+intersection.manhattanDistance(&segment1.start)+intersection.manhattanDistance(&segment2.start))
			}
			d2 += segment2.length()
		}
		d1 += segment1.length()
	}
	return min
}

func solve3b(input *os.File) {
	buf, err := ioutil.ReadAll(input)
	aocutil.Check(err)
	w1, w2 := parseWires(string(buf))
	fmt.Println(shortestDelay(w1, w2))
}
