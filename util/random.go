package util

import (
	"crypto/rand"
	"math/big"

	"github.com/google/uuid"
)

// RandomString generates a random string of length n.
func RandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

func UUID() string {
	return uuid.New().String()
}
