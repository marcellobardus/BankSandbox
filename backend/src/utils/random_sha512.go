package utils

import (
	"crypto/sha256"
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

// NewRandomSha256 return a new hex-string encoded random sha256 hash
func NewRandomSha256() string {
	randomString := uniuri.New()
	sha256Hash := sha256.Sum256([]byte(randomString))
	return hex.EncodeToString(sha256Hash[:])
}
