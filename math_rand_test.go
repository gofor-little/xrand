package xrand_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gofor-little/xrand"
)

func TestRandomMathBytes(t *testing.T) {
	var length uint = 16

	data, err := xrand.MathBytes(length)
	require.NoError(t, err)
	require.Equal(t, length, uint(len(data)))
}

func TestRandomMathString(t *testing.T) {
	var length uint = 16

	data, err := xrand.MathString(length)
	require.NoError(t, err)
	require.Equal(t, length, uint(len(data)))
}
