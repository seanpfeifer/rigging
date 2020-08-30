// Package ptr contains utility functions for dealing with pointers.
// This is particular useful in the case of assembling Protobuf structs where making a temporary var
// to do a pointer assignment decreases readability.
//
// The existence of this is necessary because eg `String() string` doesn't assign the value to an address (var) upon return,
// so `*val = getString()` tries to get the address of a thing that DOESN'T HAVE ONE - hence a "invalid memory address or nil
// pointer dereference".
//
// This was inspired by Stripe's use of similar funcs to fill Protobufs (https://github.com/stripe/stripe-go/blob/7a5549d0ab9528293ab2a8f0f1367674062246af/stripe.go#L877).
package ptr

func Bool(v bool) *bool                   { return &v }
func Int(v int) *int                      { return &v }
func Int8(v int8) *int8                   { return &v }
func Int16(v int16) *int16                { return &v }
func Int32(v int32) *int32                { return &v }
func Int64(v int64) *int64                { return &v }
func Uint(v uint) *uint                   { return &v }
func Uint8(v uint8) *uint8                { return &v }
func Uint16(v uint16) *uint16             { return &v }
func Uint32(v uint32) *uint32             { return &v }
func Uint64(v uint64) *uint64             { return &v }
func Uintptr(v uintptr) *uintptr          { return &v }
func Float32(v float32) *float32          { return &v }
func Float64(v float64) *float64          { return &v }
func Complex64(v complex64) *complex64    { return &v }
func Complex128(v complex128) *complex128 { return &v }
func String(v string) *string             { return &v }
