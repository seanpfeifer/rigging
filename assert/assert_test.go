package assert

import (
	"fmt"
	"testing"
	"time"
)

type testCase struct {
	Name          string
	Expected      interface{}
	Actual        interface{}
	Pass          bool
	loggedMessage string
	Epsilon       time.Duration // For duration tests
}

func (t *testCase) Errorf(format string, args ...interface{}) {
	t.loggedMessage = fmt.Sprintf(format, args...)
}

func TestExpectedActual(t *testing.T) {
	a, b := testCase{Name: "a"}, testCase{Name: "b"}
	var nilCase, nilCase2 *testCase
	c := []testCase{
		{Name: "strings pass", Expected: " ", Actual: " ", Pass: true},
		{Name: "strings fail", Expected: "a", Actual: "b", Pass: false},
		{Name: "struct pass", Expected: a, Actual: a, Pass: true},
		{Name: "struct fail", Expected: a, Actual: b, Pass: false},
		{Name: "pointer pass", Expected: &a, Actual: &a, Pass: true},
		{Name: "pointer fail", Expected: &a, Actual: &b, Pass: false},
		{Name: "nil pass", Expected: nilCase, Actual: nilCase2, Pass: true},
		// This is a very important test case - remember that we're not just comparing "nil", but also the type!
		{Name: "nil cast pass", Expected: nilCase, Actual: (*testCase)(nil), Pass: true},
		// In this case "nil" is not the same as a nil testCase!
		{Name: "nil fail", Expected: nilCase, Actual: nil, Pass: false},
		{Name: "int pass", Expected: 12, Actual: 12, Pass: true},
		{Name: "int fail", Expected: 12, Actual: 42, Pass: false},
		// Again, remember that types matter!
		{Name: "int fail", Expected: 12, Actual: int8(12), Pass: false},
		// ... Even if the "int" is the same size as the system's "int"s, this should STILL fail because they're not the same.
		{Name: "int fail", Expected: 12, Actual: int32(12), Pass: false},
		{Name: "int fail", Expected: 12, Actual: int64(12), Pass: false},
	}

	for i, tc := range c {
		pass := ExpectedActual(&tc, tc.Expected, tc.Actual, tc.Name)
		if pass != tc.Pass {
			t.Errorf(`[%d] %s failed: "%s"`, i, tc.Name, tc.loggedMessage)
		}
	}
}

func TestExpectedApproxTime(t *testing.T) {
	t1, t2 := time.Unix(1593934208, 12), time.Unix(1593934208, 12)
	c := []testCase{
		{Name: "same time struct", Expected: t1, Actual: t1, Pass: true},
		{Name: "same time", Expected: t1, Actual: t2, Pass: true},
		{Name: "smaller delta", Expected: t1, Actual: t1.Add(time.Millisecond * 2), Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "border delta", Expected: t1, Actual: t1.Add(time.Millisecond * 3), Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "larger delta", Expected: t1, Actual: t1.Add(time.Millisecond * 4), Pass: false, Epsilon: time.Millisecond * 3},
	}

	for i, tc := range c {
		pass := ExpectedApproxTime(&tc, tc.Expected.(time.Time), tc.Actual.(time.Time), tc.Epsilon, tc.Name)
		if pass != tc.Pass {
			t.Errorf(`[%d] %s failed: "%s"`, i, tc.Name, tc.loggedMessage)
		}
	}
}
