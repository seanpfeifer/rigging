package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// RandomID is a random 16 byte ID. This is similar to a random UUID (RFC 4122), except it lacks the formatting and is
// purely 128 random bits. This is for simple use cases where all we care about is that we have a unique, random identifier.
//
// If you want something that abides by RFC 4122, use https://github.com/google/uuid
type RandomID [16]byte

// NewRandom returns a new random ID, or an error if we fail to read from crypto/rand.
func NewRandom() (RandomID, error) {
	var id RandomID
	_, err := io.ReadFull(rand.Reader, id[:])

	return id, err
}

// String returns the RandomID as a 32-character hex string, with no separators.
func (r RandomID) String() string {
	return hex.EncodeToString(r[:])
}
