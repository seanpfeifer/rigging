// Package password contains password hashing + verification functionality.
package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashCost = 12 // OWASP recommends raising this to 12 from the default of 10
	pepperLen        = 16 // Length of the pepper we're going to use
	// Max password length to allow into the system. Empirically tested via unit test to be 72 useful characters (including pepper) with bcrypt.
	// Should be "72-pepperLen" or less for each character to actually matter.
	maxPasswordLen = 50
)

// SaltPepperHash applies a salt + pepper to the given password and returns the hash.
// The pepper is expected to be secret and stored on eg a server machine (NOT in the database)
// as another layer of security.
func SaltPepperHash(givenPass string, pepper []byte) ([]byte, error) {
	if len(givenPass) > maxPasswordLen || len(givenPass) == 0 {
		return nil, errors.New("password does not meet length requirements")
	}
	pepperedPass := append([]byte(givenPass), pepper...)

	return bcrypt.GenerateFromPassword(pepperedPass, passwordHashCost)
}

// IsValidSaltPepperHash returns true if the salt + pepper + password equals the given hash.
func IsValidSaltPepperHash(givenPass string, pepper []byte, hash []byte) bool {
	if len(givenPass) > maxPasswordLen || len(givenPass) == 0 {
		return false
	}
	pepperedPass := append([]byte(givenPass), pepper...)

	// We're OK if this returns a nil error
	return bcrypt.CompareHashAndPassword(hash, pepperedPass) == nil
}
