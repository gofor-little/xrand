package xrand

import "math/rand"

// MathBytes generates a random byte slice. This is NOT cryptographically
// secure. If you require the bytes to be cryptographically secure, see
// xrand.RandomCryptoBytes.
func MathBytes(length uint) ([]byte, error) {
	key := make([]byte, length)

	/* #nosec */
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// MathString generates a random string. This is NOT cryptographically secure.
// If you require the string to be cryptographically secure see, xrand.RandomCryptoString.
func MathString(length uint) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	bytes, err := MathBytes(length)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}
