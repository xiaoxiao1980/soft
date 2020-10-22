package util

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(s string) [32]byte {
	sum := sha256.Sum256([]byte(s))
	fmt.Printf("%x", sum)
	return sum
}
