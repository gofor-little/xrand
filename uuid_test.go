package xrand_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gofor-little/xrand"
)

func TestUUIDV4(t *testing.T) {
	uuid, err := xrand.UUIDV4()
	require.NoError(t, err)
	require.Equal(t, 36, len(uuid))
}
