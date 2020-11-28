package xrand_test

import (
	"testing"

	"github.com/gofor-little/xrand"
)

func TestRandomMathBytes(t *testing.T) {
	var length uint = 16

	data, err := xrand.RandomMathBytes(length)
	if err != nil {
		t.Fatalf("failed to generate random bytes: %v", err)
	}

	if uint(len(data)) != length {
		t.Fatalf("xrand.RandomMathBytes returned incorrect length, wanted: %d, got: %d", length, len(data))
	}
}

func TestRandomMathString(t *testing.T) {
	var length uint = 16

	data, err := xrand.RandomMathString(length)
	if err != nil {
		t.Fatalf("failed to generate random string: %v", err)
	}

	if uint(len(data)) != length {
		t.Fatalf("xrand.RandomMathString returned incorrect length, wanted: %d, got: %d", length, len(data))
	}
}
