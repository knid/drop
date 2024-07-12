package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomBytes(n int) []byte {
	token := make([]byte, n)
	rand.Read(token)
	return token
}

func RandomString(n int) string {
	return hex.EncodeToString(RandomBytes(n))
}
