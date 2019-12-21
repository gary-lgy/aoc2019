package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	solvers["14a"] = solve14a
	solvers["14b"] = solve14b
}

type reaction struct {
	reactants   []string
	reactantQty []int
	product     string
	productQty  int
}

func readReactions(input io.Reader) (map[string]*reaction, error) {
	scanner := bufio.NewScanner(input)

	reactions := make(map[string]*reaction)
	for scanner.Scan() {
		reaction := reaction{}

		text := strings.Split(scanner.Text(), " => ")

		rhs := strings.Split(text[1], " ")
		reaction.product = rhs[1]
		productQty, err := strconv.ParseInt(rhs[0], 10, 32)
		if err != nil {
			return nil, err
		}
		reaction.productQty = int(productQty)

		lhs := strings.Split(text[0], ", ")
		for _, reactant := range lhs {
			split := strings.Split(reactant, " ")
			reaction.reactants = append(reaction.reactants, split[1])
			qty, err := strconv.ParseInt(split[0], 10, 32)
			if err != nil {
				return nil, err
			}
			reaction.reactantQty = append(reaction.reactantQty, int(qty))
		}
		reactions[reaction.product] = &reaction
	}
	return reactions, nil
}

func minOreRequired(reactions map[string]*reaction, leftovers map[string]int, product string, amount int) int {
	if product == "ORE" {
		return amount
	}
	leftAmount, hasLeftover := leftovers[product]
	if hasLeftover {
		if leftAmount >= amount {
			leftovers[product] -= amount
			return 0
		} else {
			amount -= leftAmount
			leftovers[product] = 0
		}
	}
	reaction := reactions[product]
	multiplier := amount / reaction.productQty
	if multiplier*reaction.productQty < amount {
		multiplier++
		leftovers[product] = multiplier*reaction.productQty - amount
	}
	total := 0
	for i, reactant := range reaction.reactants {
		qty := reaction.reactantQty[i] * multiplier
		total += minOreRequired(reactions, leftovers, reactant, qty)
	}
	return total
}

func solve14a(input io.Reader) (string, error) {
	reactions, err := readReactions(input)
	if err != nil {
		return "", err
	}
	leftovers := make(map[string]int)
	ans := minOreRequired(reactions, leftovers, "FUEL", 1)
	return fmt.Sprint(ans), nil
}

func solve14b(input io.Reader) (string, error) {
	oreAmount := 1000000000000
	reactions, err := readReactions(input)
	if err != nil {
		return "", err
	}
	perFuel := minOreRequired(reactions, make(map[string]int), "FUEL", 1)
	lowerBound := oreAmount / perFuel
	upperBound := oreAmount
	for lowerBound < upperBound {
		middle := (lowerBound+upperBound)/2 + 1
		if minOreRequired(reactions, make(map[string]int), "FUEL", middle) <= oreAmount {
			lowerBound = middle
		} else {
			upperBound = middle - 1
		}
	}
	return fmt.Sprint(lowerBound), nil
}
