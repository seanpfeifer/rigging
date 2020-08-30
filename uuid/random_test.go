package uuid

import (
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

func TestRandom(t *testing.T) {
	id, err := NewRandom()
	ExpectedActual(t, nil, err, "new random")
	str := id.String()
	ExpectedActual(t, 32, len(str), "hex string length")
}
