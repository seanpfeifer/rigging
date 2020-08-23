package signing

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"errors"
)

// HMACKeySize is the expected size of an HMAC signing key.
const HMACKeySize = 64

// HMACKey is a secret key used for message signing using HMAC+SHA256.
//
// I've chosen to use HMAC+SHA256 instead of ed25519 due to the fact that it's more standard and mostly
// equivalent in other ways for a "secret key signing" use case. This is for when don't need a public key setup
// to allow others to verify the identity of the signer, just for us to verify that we created the message.
type HMACKey [HMACKeySize]byte

func (key *HMACKey) Sign(msg string) []byte {
	h := hmac.New(sha256.New, key[:])
	h.Write([]byte(msg))
	return h.Sum(nil)
}

func (key *HMACKey) IsValid(msg string, givenMac []byte) bool {
	h := hmac.New(sha256.New, key[:])
	h.Write([]byte(msg))
	result := h.Sum(nil)
	return hmac.Equal(givenMac, result)
}

func NewHMACKey() (HMACKey, error) {
	var key HMACKey
	n, err := rand.Read(key[:])
	if err != nil {
		return key, err
	}
	if n != HMACKeySize {
		return key, errors.New("could not read full key length from rand")
	}

	return key, nil
}
