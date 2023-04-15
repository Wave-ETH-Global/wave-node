package utils

import "crypto/rand"

// nolint:gochecknoglobals
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func SecureRandom(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	result := make([]rune, n)
	for i, v := range b {
		result[i] = letterRunes[int(v)%len(letterRunes)]
	}
	return string(result), nil
}
