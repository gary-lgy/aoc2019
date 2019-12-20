package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gary-lgy/aoc2019/aocutil"
)

func init() {
	solvers["12a"] = solve12a
	solvers["12b"] = solve12b
}

func readCoordinates(input io.Reader) ([]aocutil.IntTriple, error) {
	scanner := bufio.NewScanner(input)
	var coordinates []aocutil.IntTriple
	for scanner.Scan() {
		r := strings.NewReplacer("<", "", ">", "", "x=", "", "y=", "", "z=", "", " ", "")
		values := strings.Split(r.Replace(scanner.Text()), ",")
		x, err := strconv.ParseInt(values[0], 10, 32)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseInt(values[1], 10, 32)
		if err != nil {
			return nil, err
		}
		z, err := strconv.ParseInt(values[2], 10, 32)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, aocutil.IntTriple{X: int(x), Y: int(y), Z: int(z)})
	}
	return coordinates, nil
}

func calcNewVelocityOnSingleAxis(va, vb, pa, pb int) (int, int) {
	switch {
	case pa < pb:
		return va + 1, vb - 1
	case pa > pb:
		return va - 1, vb + 1
	default:
		return va, vb
	}
}

func applyGravity(positions, velocities []aocutil.IntTriple) {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			va, vb, pa, pb := &velocities[i], &velocities[j], &positions[i], &positions[j]
			va.X, vb.X = calcNewVelocityOnSingleAxis(va.X, vb.X, pa.X, pb.X)
			va.Y, vb.Y = calcNewVelocityOnSingleAxis(va.Y, vb.Y, pa.Y, pb.Y)
			va.Z, vb.Z = calcNewVelocityOnSingleAxis(va.Z, vb.Z, pa.Z, pb.Z)
		}
	}
}

func applyVelocity(positions, velocities []aocutil.IntTriple) {
	for i := 0; i < 4; i++ {
		positions[i].X += velocities[i].X
		positions[i].Y += velocities[i].Y
		positions[i].Z += velocities[i].Z
	}
}

func simulateMoons(initialPositions []aocutil.IntTriple, steps int) ([]aocutil.IntTriple, []aocutil.IntTriple, error) {
	positions := make([]aocutil.IntTriple, 4)
	copy(positions, initialPositions)
	velocities := make([]aocutil.IntTriple, 4)
	for i := 0; i < steps; i++ {
		applyGravity(positions, velocities)
		applyVelocity(positions, velocities)
	}
	return positions, velocities, nil
}

func calcTotalEnergy(positions, velocities []aocutil.IntTriple) int {
	totalEnergy := 0
	for i := 0; i < 4; i++ {
		v, p := &velocities[i], &positions[i]
		pe := aocutil.AbsInt(p.X) + aocutil.AbsInt(p.Y) + aocutil.AbsInt(p.Z)
		ke := aocutil.AbsInt(v.X) + aocutil.AbsInt(v.Y) + aocutil.AbsInt(v.Z)
		totalEnergy += pe * ke
	}
	return totalEnergy
}

func solve12a(input io.Reader) (string, error) {
	initialPositions, err := readCoordinates(input)
	if err != nil {
		return "", err
	}
	finalPositions, finalVelocities, err := simulateMoons(initialPositions, 1000)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(calcTotalEnergy(finalPositions, finalVelocities)), nil
}

func stepsUntilRepeat(state [8]int) int64 {
	// Use single array instead of two slices to be able to hash into a map
	// First 4 are positions, last 4 are velocities
	states := make(map[[8]int]bool)
	states[state] = true
	var steps int64 = 0
	for {
		// apply gravity
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				state[i+4], state[j+4] = calcNewVelocityOnSingleAxis(state[i+4], state[j+4], state[i], state[j])
			}
		}
		// apply velocity
		for i := 0; i < 4; i++ {
			state[i] += state[i+4]
		}
		steps++
		if _, exists := states[state]; exists {
			return steps
		} else {
			states[state] = true
		}
	}
}

func solve12b(input io.Reader) (string, error) {
	initialPositions, err := readCoordinates(input)
	if err != nil {
		return "", err
	}
	x := stepsUntilRepeat([8]int{initialPositions[0].X, initialPositions[1].X, initialPositions[2].X, initialPositions[3].X})
	y := stepsUntilRepeat([8]int{initialPositions[0].Y, initialPositions[1].Y, initialPositions[2].Y, initialPositions[3].Y})
	z := stepsUntilRepeat([8]int{initialPositions[0].Z, initialPositions[1].Z, initialPositions[2].Z, initialPositions[3].Z})
	return fmt.Sprint(aocutil.LeastCommonMultiple(aocutil.LeastCommonMultiple(x, y), z)), nil
}
