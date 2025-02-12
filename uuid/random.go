package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomID is a random 16 byte ID. This is similar to a random UUID (RFC 4122), except it lacks the formatting and is
// purely 128 random bits. This is for simple use cases where all we care about is that we have a unique, random identifier.
//
// If you want something that abides by RFC 4122, use https://github.com/google/uuid
type RandomID [16]byte

// NewRandom returns a new random ID.
// It can never error, and crashes the program irrecoverably if an error is returned, per Go 1.24 crypto/rand.Read().
// This crash will only happen on legacy Linux systems (prior to verison 3.17). See https://github.com/golang/go/issues/66821
func NewRandom() RandomID {
	var id RandomID
	// crypto/rand.Read() will ALWAYS fill the buffer and not return an error, so I'm intentionally ignoring both return values here
	_, _ = rand.Read(id[:])

	return id
}

// String returns the RandomID as a 32-character hex string, with no separators.
func (r RandomID) String() string {
	return hex.EncodeToString(r[:])
}
