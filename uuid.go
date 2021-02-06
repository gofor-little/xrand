package xrand

import (
	"crypto/rand"
	"fmt"
)

// UUIDV4 generates Version 4 UUID compliant with RFC 4122.
// See https://tools.ietf.org/html/rfc4122 for more info.
// Returns an error if the systems cryptographically secure random number generator fails,
// in which case the caller should not continue.
func UUIDV4() (string, error) {
	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// Version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
