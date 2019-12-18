package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossiblePasswordA(t *testing.T) {
	assert.True(t, isPossiblePartA(111111))
	assert.False(t, isPossiblePartA(223450))
	assert.False(t, isPossiblePartA(123789))
}

func TestIsPossiblePasswordB(t *testing.T) {
	assert.True(t, isPossiblePartB(111122))
	assert.False(t, isPossiblePartB(123444))
	assert.True(t, isPossiblePartB(112233))
}
