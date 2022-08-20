// Package hashing contains simple hashing functionality. This is NOT meant for passwords, but
// rather for simply verifying the integrity of a message using a secret.
//
// Note this does not allow for any other party to verify authenticity of a message, as the
// secret should not be shared.
package hashing

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
)

// HMACKeySize is the expected size of an HMAC hashing key.
const HMACKeySize = 64

// HMACKey is a secret key used for message hashing using HMAC+SHA256.
//
// I've chosen to use HMAC+SHA256 instead of ed25519 due to the fact that it's more standard and mostly
// equivalent in other ways for a "secret key hashing" use case. This is for when don't need a public key setup
// to allow others to verify the identity of the hasher, just for us to verify that we created the message.
type HMACKey [HMACKeySize]byte

// Hash returns the SHA256 hash of the bytes in the given message, calculated using our key.
func (key *HMACKey) Hash(msg string) []byte {
	h := hmac.New(sha256.New, key[:])
	h.Write([]byte(msg))
	return h.Sum(nil)
}

// IsValid returns true if the givenMac bytes are a valid hash of the bytes in msg, calculated using our key.
func (key *HMACKey) IsValid(msg string, givenMac []byte) bool {
	h := hmac.New(sha256.New, key[:])
	h.Write([]byte(msg))
	result := h.Sum(nil)
	return hmac.Equal(givenMac, result)
}

// NewHMACKey creates a new cryptographically random HMAC key with HMACKeySize bytes.
//
// You typically will want to store the output of this and use it repeatedly, hashing messages that you send out
// and checking validity when returned to you.
func NewHMACKey() (HMACKey, error) {
	var key HMACKey
	if _, err := rand.Read(key[:]); err != nil {
		return key, err
	}

	return key, nil
}
