package gmath

import (
	"testing"

	. "github.com/seanpfeifer/rigging/assert"
)

func TestClamp(t *testing.T) {
	ExpectedActual(t, 0, Clamp(-1, 0, 100), "clamp to min (int)")
	ExpectedActual(t, 100, Clamp(101, 0, 100), "clamp to max (int)")
	ExpectedActual(t, 0.0, Clamp(-1.0, 0.0, 1.0), "clamp to min (float)")
	ExpectedActual(t, 1.0, Clamp(2.0, 0.0, 1.0), "clamp to max (float)")
}

func TestLerp(t *testing.T) {
	ExpectedActual(t, 25, Lerp(int64(0), 100, 0.25), "int64 25%")
	ExpectedActual(t, 20, Lerp(int64(20), 100, 0.0), "int64 min")
	ExpectedActual(t, 100, Lerp(int64(20), 100, 1.0), "int64 max")

	ExpectedActual(t, 75, Lerp(int64(100), 0, 0.25), "int64 reverse 25%")
	ExpectedActual(t, 100, Lerp(int64(100), 20, 0.0), "int64 reverse min")
	ExpectedActual(t, 20, Lerp(int64(100), 20, 1.0), "int64 reverse max")

	// Unsigned is generally a bad idea. Cast to a float instead.
	ExpectedActual(t, 100, Lerp(uint8(0), 100, 1.0), "uint8 max")
	ExpectedActual(t, 20, Lerp(uint8(100), 20, 1.0), "uint8 reverse max")
	// This is a problem when a > b. Instead, cast these to a float!
	// FAIL: ExpectedActual(t, 75, Lerp(uint8(100), 0, 0.25), "uint8 reverse 25%") -> ACTUAL: 139
	ExpectedActual(t, 75, Lerp(float64(uint8(100)), 0, 0.25), "uint8 reverse 25%")
}

// Note: You can run these benchmarks with a command like:
//    go test -benchtime=20000000000x -benchmem -bench .

var resultF64 float64

func BenchmarkLerpFloat64(b *testing.B) {
	var res float64
	for n := 0; n < b.N; n++ {
		res = Lerp(0.0, 1.0, 0.5)
	}
	resultF64 = res
}

var resultF32 float32

func BenchmarkLerpFloat32(b *testing.B) {
	var res float32
	for n := 0; n < b.N; n++ {
		// Explicitly using float32 here as a convenience so we don't have to cast the params.
		// In a real-world case you'd just pass your float32 vars / constants.
		res = Lerp[float32](0.0, 1.0, 0.5)
	}
	resultF32 = res
}

var resultI64 int64

func BenchmarkLerpInt64(b *testing.B) {
	var res int64
	for n := 0; n < b.N; n++ {
		res = Lerp[int64](0, 100, 0.5)
	}
	resultI64 = res
}

var resultI32 int32

func BenchmarkLerpInt32(b *testing.B) {
	var res int32
	for n := 0; n < b.N; n++ {
		res = Lerp[int32](0, 100, 0.5)
	}
	resultI32 = res
}
