package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test8a(t *testing.T) {
	input, err := os.Open(filepath.Join("input", "8.txt"))
	require.NoError(t, err)
	ans, err := solve8a(input)
	require.NoError(t, err)
	assert.Equal(t, "2562", ans)
}
