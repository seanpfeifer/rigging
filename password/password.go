// Package password contains password hashing + verification functionality.
package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashCost = 12 // OWASP recommends raising this to 12 from the default of 10
	pepperLen        = 16 // Length of the pepper we're going to use
	// Suggested max password length. Empirically tested via unit test to be 72 useful bytes (including pepper) with bcrypt.
	// Should be "72-pepperLen" or less for each byte to actually matter. Ideally at least SOME of the pepper is included
	// in the password, so try to use less than this number of bytes.
	MaxPasswordLen = 72
)

var ErrBadPasswordLength = errors.New("password does not meet length requirements")

// SaltPepperHash applies a salt + pepper to the given password and returns the hash.
// The pepper is expected to be secret and stored on eg a server machine (NOT in the database)
// as another layer of security.
// Checks against unreasonable length passwords are expected to be done before calling this.
func SaltPepperHash(givenPass string, pepper []byte) ([]byte, error) {
	if len(givenPass) == 0 {
		return nil, ErrBadPasswordLength
	}
	pepperedPass := append([]byte(givenPass), pepper...)

	return bcrypt.GenerateFromPassword(pepperedPass, passwordHashCost)
}

// IsValidSaltPepperHash returns true if the salt + pepper + password equals the given hash.
// Checks against unreasonable length passwords are expected to be done before calling this.
func IsValidSaltPepperHash(givenPass string, pepper []byte, hash []byte) bool {
	if len(givenPass) == 0 {
		return false
	}
	pepperedPass := append([]byte(givenPass), pepper...)

	// We're OK if this returns a nil error
	return bcrypt.CompareHashAndPassword(hash, pepperedPass) == nil
}
