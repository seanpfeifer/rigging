// Package password contains password hashing + verification functionality.
// TLDR; "just use bcrypt".
//
// bcrypt is salting the password as part of hashing. If you're sure you want to pepper,
// do it to the password prior to passing here.
package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashCost = 12 // OWASP recommends raising this to 12 from the default of 10
	// Max password length in bytes. bcrypt only takes into account 72 bytes, and will error if given more.
	MaxPasswordLen = 72
)

var ErrPasswordEmpty = errors.New("password is empty")
var ErrPasswordTooLong = errors.New("password length exceeds 72 bytes")

// Hash applies a salt to the given password and returns the hash.
// This will return ErrPasswordTooLong if the password is longer than MaxPasswordLen,
// and ErrPasswordEmpty if the password is empty.
func Hash(givenPass string) ([]byte, error) {
	if len(givenPass) == 0 {
		return nil, ErrPasswordEmpty
	}
	if len(givenPass) > MaxPasswordLen {
		return nil, ErrPasswordTooLong
	}

	return bcrypt.GenerateFromPassword([]byte(givenPass), passwordHashCost)
}

// IsValid returns true if the hashed password equals the given hash.
// Checks against unreasonable length passwords are expected to be done before calling this.
// For backwards compatibility, this will not error if the given password is too long, and
// only operates on the first 72 bytes of givenPass.
func IsValid(hash []byte, givenPass string) bool {
	if len(givenPass) == 0 {
		return false
	}

	// We're OK if this returns a nil error
	return bcrypt.CompareHashAndPassword(hash, []byte(givenPass)) == nil
}
