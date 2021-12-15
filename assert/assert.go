// Package assert provides a very simple set of functions to improve readability of tests.
// Note that this is simply using equality to check the values passed in, and therefore you must be mindful of
// any necessary casts to make your types match. eg, "int" vs "int32" types matter.
package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

// The idea here is to accept `*testing.T`, or another impl to be able to test this package itself.
type Tester interface {
	Errorf(format string, args ...any)
}

// ExpectedActual logs a testing error and returns false if the expected and actual values are not equal.
// Typically you will not need the return value unless you want to stop testing on failure.
func ExpectedActual[T comparable](t Tester, expected, actual T, name string) bool {
	if expected == actual {
		return true
	}

	errStr := fmt.Sprintf(`[%s] Expected: "%+v".(%s)  Actual: "%+v".(%s)`, name,
		expected, reflect.TypeOf(expected),
		actual, reflect.TypeOf(actual))

	logErrorMsg(t, errStr)
	return false
}

// ExpectedApproxTime logs a testing error and returns false if the expected and actual time values are not close.
// Specifically, if the absolute value of the difference between the two is LARGER than the given "epsilon".
// Typically you will not need the return value unless you want to stop testing on failure.
func ExpectedApproxTime(t Tester, expected, actual time.Time, epsilon time.Duration, name string) bool {
	delta := abs(actual.Sub(expected))
	if delta > epsilon {
		errStr := fmt.Sprintf(`[%s] Expected: "%s"  Actual: "%s" Delta: %s Tolerance: %s`, name, expected, actual, delta, epsilon)

		logErrorMsg(t, errStr)
		return false
	}
	return true
}

// ExpectedApproxDuration logs a testing error and returns false if the expected and actual duration values are not close.
// Specifically, if the absolute value of the difference between the two is LARGER than the given "epsilon".
// Typically you will not need the return value unless you want to stop testing on failure.
func ExpectedApproxDuration(t Tester, expected, actual time.Duration, epsilon time.Duration, name string) bool {
	delta := abs(actual - expected)
	if delta > epsilon {
		errStr := fmt.Sprintf(`[%s] Expected: "%s"  Actual: "%s" Delta: %s Tolerance: %s`, name, expected, actual, delta, epsilon)

		logErrorMsg(t, errStr)
		return false
	}
	return true
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

func logErrorMsg(t Tester, errStr string) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		file = filepath.Base(file)
	} else {
		file = "???"
		line = 1
	}
	t.Errorf("%s:%d %s", file, line, errStr)
}
