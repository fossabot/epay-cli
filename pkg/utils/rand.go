package utils

import "math/rand"

func RandString(n int) string {
	const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = alphabets[rand.Intn(len(alphabets))]
	}

	return string(b)
}
