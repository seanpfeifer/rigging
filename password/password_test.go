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

	// Change the last byte in the password and make sure the password fails.
	password[MaxPasswordLen-1] = ^password[MaxPasswordLen-1]
	isValid = IsValidSaltPepperHash(string(password[:]), pepper[:], hashed)
	ExpectedActual(t, false, isValid, "verifying bad hash")

	var longPassword [MaxPasswordLen * 2]byte
	hashed, err = SaltPepperHash(string(longPassword[:]), pepper[:])
	ExpectedActual(t, nil, err, "hashing long password")

	isValid = IsValidSaltPepperHash(string(longPassword[:]), pepper[:], hashed)
	ExpectedActual(t, true, isValid, "verifying long password")
}

func TestBadPasswordHashing(t *testing.T) {
	var pepper [pepperLen]byte
	rand.Read(pepper[:])

	// Expect a failure if the password is empty
	_, err := SaltPepperHash("", pepper[:])
	ExpectedActual(t, ErrBadPasswordLength, err, "expecting error for zero-length password")

	hashed, err := SaltPepperHash("random password here", pepper[:])
	ExpectedActual(t, nil, err, "hashing password")
	// Now check inputs on validation.
	// Zero length
	isValid := IsValidSaltPepperHash("", pepper[:], hashed)
	ExpectedActual(t, false, isValid, "validating bad empty password")
}
