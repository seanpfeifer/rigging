package password

import (
	"crypto/rand"
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

func TestPasswordHashing(t *testing.T) {
	var pepper [pepperLen]byte
	rand.Read(pepper[:])

	var password [MaxPasswordLen]byte
	rand.Read(password[:])

	hashed, err := SaltPepperHash(string(password[:]), pepper[:])
	ExpectedActual(t, nil, err, "hashing password")

	isValid := IsValidSaltPepperHash(string(password[:]), pepper[:], hashed)
	ExpectedActual(t, true, isValid, "verifying hash")

	// Change the last byte in the pepper and make sure the password fails.
	// This was used to empirically test that the max (password + pepper) length for bcrypt is 72.
	pepper[pepperLen-1] = ^pepper[pepperLen-1]
	isValid = IsValidSaltPepperHash(string(password[:]), pepper[:], hashed)
	ExpectedActual(t, false, isValid, "verifying bad hash")
}

func TestBadPasswordHashing(t *testing.T) {
	var pepper [pepperLen]byte
	rand.Read(pepper[:])

	// Expect a failure if the password is too long for both hashing + validating
	var badPassword [MaxPasswordLen + 1]byte
	_, err := SaltPepperHash(string(badPassword[:]), pepper[:])
	// Expect a non-nil error (specifically ErrBadPasswordLength)
	ExpectedActual(t, ErrBadPasswordLength, err, "expecting error for overly long password")

	// Same with a zero-len password
	_, err = SaltPepperHash("", pepper[:])
	ExpectedActual(t, ErrBadPasswordLength, err, "expecting error for zero-length password")

	hashed, err := SaltPepperHash("random password here", pepper[:])
	ExpectedActual(t, nil, err, "hashing password")
	// Now check inputs on validation.
	// Too long
	isValid := IsValidSaltPepperHash(string(badPassword[:]), pepper[:], hashed)
	ExpectedActual(t, false, isValid, "validating bad long password")

	// Too short
	isValid = IsValidSaltPepperHash("", pepper[:], hashed)
	ExpectedActual(t, false, isValid, "validating bad short password")
}
