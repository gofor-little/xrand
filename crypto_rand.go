package xrand

import "crypto/rand"

// RandomCryptoBytes generates a cryptographically secure random byte slice.
// Returns an error if the systems cryptographically secure random generator fails,
// in which case the caller should not continue.
func RandomCryptoBytes(length uint) ([]byte, error) {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// RandomCryptoString generates a cryptographically secure random string.
// Returns an error if the systems cryptographically secure random generator fails,
// in which case the caller should not continue.
func RandomCryptoString(length uint) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	bytes, err := RandomCryptoBytes(length)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}
