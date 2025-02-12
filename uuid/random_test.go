package uuid

import (
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

func TestRandom(t *testing.T) {
	id := NewRandom()
	str := id.String()
	ExpectedActual(t, 32, len(str), "hex string length")
}
