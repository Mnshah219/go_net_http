package utils

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

type Argon2idHash struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

// NewArgon2idHash constructor function for
// Argon2idHash.
func NewArgon2idHash() Argon2idHash {
	var time uint32 = 1
	var saltLen uint32 = 32
	var memory uint32 = 64 * 1024
	var threads uint8 = 1
	var keyLen uint32 = 32
	return Argon2idHash{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

// GenerateHash using the password and provided salt.
// If not salt value provided fallback to random value
// generated of a given length.
func GenerateHash(password, salt []byte) string {
	argon2idHash := NewArgon2idHash()
	// Generate hash
	hash := argon2.IDKey(password, salt, argon2idHash.time, argon2idHash.memory, argon2idHash.threads, argon2idHash.keyLen)
	// Return the generated hash and salt used for storage.
	return base64.StdEncoding.EncodeToString(hash)
}

// Compare generated hash with store hash.
func Compare(expectedHash string, salt, password []byte) error {
	// Generate hash for comparison.
	generatedHash := GenerateHash(password, salt)
	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if generatedHash != expectedHash {
		return errors.New("hash doesn't match")
	}
	return nil
}
