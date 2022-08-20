// Package ptr contains utility functions for dealing with pointers.
// This is particular useful in the case of assembling Protobuf structs where making a temporary var
// to do a pointer assignment decreases readability.
//
// The existence of this is necessary because eg `String() string` doesn't assign the value to an address (var) upon return, so
//
//	*val = getString()
//
// tries to get the address of a thing that DOESN'T HAVE ONE - hence an "invalid memory address or nil
// pointer dereference".
// A good example is
//
//	item{Name: &"test"}
//
// does not compile, so use
//
//	item{Name: ptr.Ptr("test")}
//
// See unit tests for a demonstration.
//
// This was inspired by Stripe's use of similar funcs to fill Protobufs (https://github.com/stripe/stripe-go/blob/7a5549d0ab9528293ab2a8f0f1367674062246af/stripe.go#L877).
package ptr

// Ptr returns a type-safe reference to a variable with the given value.
// This should be used only in cases where you have to assign a pointer to a reference of a primitive literal,
// and want a simple one-liner (eg assigning to struct members).
//
// If you can simply do
// 	x := &v
// then do that instead of using this func.
func Ptr[T any](v T) *T { return &v }
