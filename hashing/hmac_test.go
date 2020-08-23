package hashing

import (
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

const (
	dataToBeHashed   = "My Account Info Here"
	expectedHashSize = 32
)

func TestIsValid(t *testing.T) {
	key, err := NewHMACKey()
	if err != nil {
		t.FailNow()
	}

	// Reset the timer, since we don't want to time the setup we had to do
	hash := key.Hash(dataToBeHashed)
	ExpectedActual(t, expectedHashSize, len(hash), "hash not expected length")

	verified := key.IsValid(dataToBeHashed, hash)
	ExpectedActual(t, true, verified, "matching hashes")

	verified = key.IsValid(dataToBeHashed+"1", hash)
	ExpectedActual(t, false, verified, "differing hashes")
}

func BenchmarkHashHMAC(b *testing.B) {
	key, err := NewHMACKey()
	if err != nil {
		b.FailNow()
	}

	// Reset the timer, since we don't want to time the setup we had to do
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key.Hash(dataToBeHashed)
	}
}

func BenchmarkVerifyHMAC(b *testing.B) {
	key, err := NewHMACKey()
	if err != nil {
		b.FailNow()
	}
	hash := key.Hash(dataToBeHashed)

	// Reset the timer, since we don't want to time the setup we had to do
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key.IsValid(dataToBeHashed, hash)
	}
}
