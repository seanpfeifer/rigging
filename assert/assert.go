// Package assert provides a very simple set of functions to improve readability of tests.
// Note that this is simply using equality to check the values passed in, and therefore you must be mindful of
// any necessary casts to make your types match. eg, "int" vs "int32" types matter.
package assert

import (
	"reflect"
	"time"

	"github.com/seanpfeifer/rigging/num"
)

// The idea here is to accept `*testing.T`, or another impl to be able to test this package itself.
type Tester interface {
	Errorf(format string, args ...any)
	Helper()
}

// ExpectedActual logs a testing error and returns false if the expected and actual values are not equal.
// Typically you will not need the return value unless you want to stop testing on failure.
// If you call this within a test utility func, make sure you use `t.Helper()` so you get accurate failure locations.
func ExpectedActual[V any](t Tester, expected, actual V, name string) bool {
	if reflect.DeepEqual(expected, actual) {
		return true
	}

	t.Helper() // Marks this func as a Helper, so this error gets logged at the caller's location
	t.Errorf(`[%s] Expected: "%+v".(%s)  Actual: "%+v".(%s)`, name,
		expected, reflect.TypeOf(expected),
		actual, reflect.TypeOf(actual))

	return false
}

// ExpectedApproxTime logs a testing error and returns false if the expected and actual time values are not close.
// Specifically, if the absolute value of the difference between the two is LARGER than the given "epsilon".
// Typically you will not need the return value unless you want to stop testing on failure.
func ExpectedApproxTime(t Tester, expected, actual time.Time, epsilon time.Duration, name string) bool {
	delta := absDuration(actual.Sub(expected))
	if delta > epsilon {
		t.Helper() // Marks this func as a Helper, so this error gets logged at the caller's location
		t.Errorf(`[%s] Expected: "%s"  Actual: "%s" Delta: %s Tolerance: %s`, name, expected, actual, delta, epsilon)
		return false
	}
	return true
}

// ExpectedApprox logs a testing error and returns false if the expected and actual values are not close.
// Specifically, if the absolute value of the difference between the two is LARGER than the given "epsilon".
// Typically you will not need the return value unless you want to stop testing on failure.
func ExpectedApprox[V num.Real](t Tester, expected, actual, epsilon V, name string) bool {
	delta := getDelta(actual, expected)
	if delta > epsilon {
		t.Helper() // Marks this func as a Helper, so this error gets logged at the caller's location
		t.Errorf(`[%s] Expected: "%s"  Actual: "%s" Delta: %s Tolerance: %s`, name, expected, actual, delta, epsilon)
		return false
	}
	return true
}

// This exists because subtractions for unsigned numbers cannot possibly go negative, so I want to just
// avoid doing the subtraction. Since this is for testing, relative performance for this isn't a big concern.
func getDelta[V num.Real](a, b V) V {
	if a > b {
		return a - b
	}
	return b - a
}

// This is explicitly to support time.Time delta
func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
