package xrand_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gofor-little/xrand"
)

func TestCryptoBytes(t *testing.T) {
	var length uint = 16

	data, err := xrand.CryptoBytes(length)
	require.NoError(t, err)
	require.Equal(t, length, uint(len(data)))
}

func TestCryptoString(t *testing.T) {
	var length uint = 16

	data, err := xrand.CryptoString(length)
	require.NoError(t, err)
	require.Equal(t, length, uint(len(data)))
}
