package password

import (
	"crypto/rand"
	"testing"
)

func TestPasswordHashing(t *testing.T) {
	const pepperLen = 16
	var pepper [pepperLen]byte
	rand.Read(pepper[:])

	var password [maxPasswordLen]byte
	rand.Read(password[:])

	hashed, err := SaltPepperHash(string(password[:]), pepper[:])
	if err != nil {
		t.Fail()
	}

	if !IsValidSaltPepperHash(string(password[:]), pepper[:], hashed) {
		t.Fail()
	}

	// Change the last byte in the pepper and make sure the password fails.
	// This was used to empirically test that the max (password + pepper) length for bcrypt is 72.
	pepper[pepperLen-1] = ^pepper[pepperLen-1]
	if IsValidSaltPepperHash(string(password[:]), pepper[:], hashed) {
		t.Fail()
	}
}

func TestBadPasswordHashing(t *testing.T) {
	const pepperLen = 16
	var pepper [pepperLen]byte
	rand.Read(pepper[:])

	// Expect a failure if the password is too long for both hashing + validating
	var badPassword [maxPasswordLen + 1]byte
	_, err := SaltPepperHash(string(badPassword[:]), pepper[:])
	// Expect a non-nil error
	if err == nil {
		t.Fail()
	}

	// Same with a zero-len password
	_, err = SaltPepperHash("", pepper[:])
	// Expect a non-nil error
	if err == nil {
		t.Fail()
	}

	hashed, err := SaltPepperHash("random password here", pepper[:])
	if err != nil {
		t.Fail()
	}
	// Now check inputs on validation.
	// Too long
	if IsValidSaltPepperHash(string(badPassword[:]), pepper[:], hashed) {
		t.Fail()
	}
	// Too short
	if IsValidSaltPepperHash("", pepper[:], hashed) {
		t.Fail()
	}
}
