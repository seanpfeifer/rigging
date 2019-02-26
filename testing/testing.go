package testing

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

// ExpectedActual logs a testing error if the expected and actual values are not equal.
func ExpectedActual(t *testing.T, expected, actual interface{}, name string) {
	if expected != actual {
		errStr := fmt.Sprintf(`[%s] Expected: "%+v".(%s)  Actual: "%+v".(%s)`, name,
			expected, reflect.TypeOf(expected),
			actual, reflect.TypeOf(actual))

		logErrorMsg(t, errStr)
	}
}

// ExpectedApproxTime logs a testing error if the expected and actual time values are not close.
func ExpectedApproxTime(t *testing.T, expected, actual time.Time, epsilon time.Duration, name string) {
	deltaTime := actual.Sub(expected)
	if deltaTime > epsilon {
		errStr := fmt.Sprintf(`[%s] Expected: "%s"  Actual: "%s" Delta: %s Tolerance: %s`, name, expected, actual, deltaTime, epsilon)

		logErrorMsg(t, errStr)
	}
}

func logErrorMsg(t *testing.T, errStr string) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, `\`); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	t.Errorf("%s:%d %s", file, line, errStr)
}
