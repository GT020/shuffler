package utils

import (
	"crypto/rand"
	"math/big"
)

func SecureRandomInt(min int, max int) int {

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))

	if err != nil {
		panic(err)
	}

	return int(n.Int64()) + min
}
