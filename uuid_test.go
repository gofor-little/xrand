package xrand_test

import (
	"testing"

	"github.com/gofor-little/xrand"
)

func TestGenerateUUID(t *testing.T) {
	uuid, err := xrand.GenerateUUID()
	if err != nil {
		t.Fatalf("failed to generate v4 UUID: %v", err)
	}

	if len(uuid) != 36 {
		t.Fatalf("xrand.GenerateUUID returned incorrect length, wanted: %d, got: %d", 36, len(uuid))
	}
}
