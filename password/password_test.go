package password

import (
	"crypto/rand"
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

func TestPasswordHashing(t *testing.T) {
	var password [MaxPasswordLen]byte
	rand.Read(password[:])

	hashed, err := Hash(string(password[:]))
	ExpectedActual(t, nil, err, "hashing password")

	isValid := IsValid(hashed, string(password[:]))
	ExpectedActual(t, true, isValid, "verifying hash")

	// Change the last byte in the password and make sure the password fails.
	password[MaxPasswordLen-1] = ^password[MaxPasswordLen-1]
	isValid = IsValid(hashed, string(password[:]))
	ExpectedActual(t, false, isValid, "verifying bad hash")

	var longPassword [MaxPasswordLen * 2]byte
	hashed, err = Hash(string(longPassword[:]))
	ExpectedActual(t, nil, err, "hashing long password")

	isValid = IsValid(hashed, string(longPassword[:]))
	ExpectedActual(t, true, isValid, "verifying long password")
}

func TestBadPasswordHashing(t *testing.T) {
	// Expect a failure if the password is empty
	_, err := Hash("")
	ExpectedActual(t, ErrEmptyPassword, err, "expecting error for zero-length password")

	hashed, err := Hash("random password here")
	ExpectedActual(t, nil, err, "hashing password")
	// Now check inputs on validation.
	// Zero length
	isValid := IsValid(hashed, "")
	ExpectedActual(t, false, isValid, "validating bad empty password")
}
