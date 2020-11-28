package xrand_test

import (
	"testing"

	"github.com/gofor-little/xrand"
)

func TestRandomCryptoBytes(t *testing.T) {
	var length uint = 16

	data, err := xrand.RandomCryptoBytes(length)
	if err != nil {
		t.Fatalf("failed to generate cryptographically secure random bytes: %v", err)
	}

	if uint(len(data)) != length {
		t.Fatalf("xrand.RandomCryptoBytes returned incorrect length, wanted: %d, got: %d", length, len(data))
	}
}

func TestRandomCryptoString(t *testing.T) {
	var length uint = 16

	data, err := xrand.RandomCryptoString(length)
	if err != nil {
		t.Fatalf("failed to generate cryptographically secure random string: %v", err)
	}

	if uint(len(data)) != length {
		t.Fatalf("xrand.RandomCryptoString returned incorrect length, wanted: %d, got: %d", length, len(data))
	}
}
