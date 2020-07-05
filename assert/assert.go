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
	Errorf(format string, args ...interface{})
}

// ExpectedActual logs a testing error and returns false if the expected and actual values are not equal.
// Typically you will not need the return value unless you want to stop testing.
func ExpectedActual(t Tester, expected, actual interface{}, name string) bool {
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
// Specifically, if the difference between the two is LARGER than the given epsilon.
// Typically you will not need the return value unless you want to stop testing.
func ExpectedApproxTime(t Tester, expected, actual time.Time, epsilon time.Duration, name string) bool {
	deltaTime := actual.Sub(expected)
	if deltaTime > epsilon {
		errStr := fmt.Sprintf(`[%s] Expected: "%s"  Actual: "%s" Delta: %s Tolerance: %s`, name, expected, actual, deltaTime, epsilon)

		logErrorMsg(t, errStr)
		return false
	}
	return true
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
