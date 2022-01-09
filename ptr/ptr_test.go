package ptr

import (
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

type item struct {
	Name   *string
	Age    *uint64
	Weight *float64
}

func TestBasic(t *testing.T) {
	// This is what we'd have to do without a function like the one in this package.
	// This is one extra explicit variable creation per attribute of our struct.
	name := "test"
	age := uint64(123)
	verboseItem := item{
		Name: &name,
		Age:  &age,
	}
	// ... imagine we have to do some calculations ...
	// *verboseItem.Weight = 42.314 // Can't do this! This will panic since Weight is still nil
	weight := 42.314
	verboseItem.Weight = &weight
	ExpectedActual(t, "test", *verboseItem.Name, "name value")
	ExpectedActual(t, 123, *verboseItem.Age, "age value")
	ExpectedActual(t, 42.314, *verboseItem.Weight, "weight value")

	// Here's the func in use
	conciseItem := item{
		Name: Ptr("test"),
		Age:  Ptr(uint64(123)),
	}
	// ... imagine we have to do some calculations ...
	conciseItem.Weight = Ptr(42.314)
	ExpectedActual(t, "test", *conciseItem.Name, "name value")
	ExpectedActual(t, 123, *conciseItem.Age, "age value")
	ExpectedActual(t, 42.314, *conciseItem.Weight, "weight value")
}
