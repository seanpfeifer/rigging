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
	_, err = Hash(string(longPassword[:]))
	// Crypto now returns an error if the password is too long, so expect that.
	ExpectedActual(t, ErrPasswordTooLong, err, "hashing long password")

	// This is the truncation that was previously done in bcrypt, useful for testing the validation below.
	hashed, err = Hash(string(longPassword[:MaxPasswordLen]))
	ExpectedActual(t, nil, err, "hashing truncated password")

	// We should still be able to validate the password if it's too long.
	isValid = IsValid(hashed, string(longPassword[:]))
	ExpectedActual(t, true, isValid, "verifying long password")
}

func TestBadPasswordHashing(t *testing.T) {
	// Expect a failure if the password is empty
	_, err := Hash("")
	ExpectedActual(t, ErrPasswordEmpty, err, "expecting error for zero-length password")

	hashed, err := Hash("random password here")
	ExpectedActual(t, nil, err, "hashing password")
	// Now check inputs on validation.
	// Zero length
	isValid := IsValid(hashed, "")
	ExpectedActual(t, false, isValid, "validating bad empty password")
}
