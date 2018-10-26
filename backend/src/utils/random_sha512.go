package utils

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/dchest/uniuri"
)

// NewRandomSha512 return a new hex-string encoded random sha512 hash
func NewRandomSha512() string {
	randomString := uniuri.New()
	sha512Hash := sha512.Sum512([]byte(randomString))
	return hex.EncodeToString(sha512Hash[:])
}
