package xrand

import "math/rand"

// RandomMathBytes generates a random byte slice. This is NOT cryptographically
// secure. If you require the bytes to be cryptographically secure, see
// xrand.RandomCryptoBytes.
func RandomMathBytes(length uint) ([]byte, error) {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// RandomMathString generates a random string. This is NOT cryptographically secure.
// If you require the string to be cryptographically secure see, xrand.RandomCryptoString.
func RandomMathString(length uint) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	bytes, err := RandomMathBytes(length)
	if err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}
