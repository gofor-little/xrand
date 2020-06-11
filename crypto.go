package rand

import "crypto/rand"

// GenerateCryptoBytes generates a cryptographically secure byte slice.
// Returns an error if the systems cryptographically secure random number generator fails,
// in which case the caller should not continue.
func GenerateCryptoBytes(length uint) ([]byte, error) {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// GenerateCryptoString generates a cryptographically secure byte slice.
// Returns an error if the systems cryptographically secure random number generator fails,
// in which case the caller should not continue.
func GenerateCryptoString(length uint) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	bytes, err := GenerateCryptoBytes(length)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}
