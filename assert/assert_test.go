package assert

import (
	"fmt"
	"testing"
	"time"

	"github.com/seanpfeifer/rigging/num"
)

type testCase struct {
	Name          string
	Expected      any
	Actual        any
	Pass          bool
	loggedMessage string
}

func (t *testCase) Errorf(format string, args ...any) {
	t.loggedMessage = fmt.Sprintf(format, args...)
}

func (t *testCase) Helper() {}

type person struct {
	Name string
	Hat  bool
}

// wrongPerson has the same contents as person, but is a different type.
// These are expected to be not equal when compared.
type wrongPerson struct {
	Name string
	Hat  bool
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
		{Name: "int8 fail", Expected: 12, Actual: int8(12), Pass: false},
		// ... Even if the "int" is the same size as the system's "int"s, this should STILL fail because they're not the same.
		{Name: "int32 fail", Expected: int(12), Actual: int32(12), Pass: false},
		{Name: "int64 fail", Expected: int(12), Actual: int64(12), Pass: false},
		// These are actually the same type, though "12" seems untyped
		{Name: "int type pass", Expected: 12, Actual: int(12), Pass: true},
		{Name: "struct pass", Expected: person{"Sean", true}, Actual: person{"Sean", true}, Pass: true},
		{Name: "struct fail", Expected: person{"Sean", true}, Actual: person{"Sean", false}, Pass: false},
		// Note that in typical usage this case would fail to even compile due to the types being different, but since we're using
		// "any" in this slice, this will compile and should fail during test.
		// eg,
		//  ExpectedActual(t, person{"Sean", true}, wrongPerson{"Sean", true}, "type test") // fails to compile due to type differences
		{Name: "struct type", Expected: person{"Sean", true}, Actual: wrongPerson{"Sean", true}, Pass: false},
	}

	for i, tc := range c {
		pass := ExpectedActual(&tc, tc.Expected, tc.Actual, tc.Name)
		if pass != tc.Pass {
			t.Errorf(`[%d] %s failed: "%s"`, i, tc.Name, tc.loggedMessage)
		}
	}
}

type testCaseApprox[T any, V num.Integer | num.Float] struct {
	Name          string
	Expected      T
	Actual        T
	Pass          bool
	loggedMessage string
	Epsilon       V
}

func (t *testCaseApprox[T, V]) Errorf(format string, args ...any) {
	t.loggedMessage = fmt.Sprintf(format, args...)
}

func (t *testCaseApprox[T, V]) Helper() {}

func TestExpectedApproxTime(t *testing.T) {
	t1, t2 := time.Unix(1593934208, 12), time.Unix(1593934208, 12)
	c := []testCaseApprox[time.Time, time.Duration]{
		// Exactness, with 0 epsilon
		{Name: "same time struct", Expected: t1, Actual: t1, Pass: true},
		{Name: "same time", Expected: t1, Actual: t2, Pass: true},
		// Test cases for the actual time being in the future
		{Name: "future smaller delta", Expected: t1, Actual: t1.Add(time.Millisecond * 2), Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "future border delta", Expected: t1, Actual: t1.Add(time.Millisecond * 3), Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "future larger delta", Expected: t1, Actual: t1.Add(time.Millisecond * 4), Pass: false, Epsilon: time.Millisecond * 3},
		// Test cases for the actual time being in the past
		{Name: "past smaller delta", Expected: t1, Actual: t1.Add(-time.Millisecond * 2), Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "past border delta", Expected: t1, Actual: t1.Add(-time.Millisecond * 3), Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "past larger delta", Expected: t1, Actual: t1.Add(-time.Millisecond * 4), Pass: false, Epsilon: time.Millisecond * 3},
	}

	for i, tc := range c {
		pass := ExpectedApproxTime(&tc, tc.Expected, tc.Actual, tc.Epsilon, tc.Name)
		if pass != tc.Pass {
			t.Errorf(`[%d] %s failed: "%s"`, i, tc.Name, tc.loggedMessage)
		}
	}
}

func TestExpectedApproxDuration(t *testing.T) {
	d := time.Millisecond * 12
	c := []testCaseApprox[time.Duration, time.Duration]{
		// Exactness, with 0 epsilon
		{Name: "same duration", Expected: d, Actual: time.Millisecond * 12, Pass: true},
		// Test cases for the actual time being in the future
		{Name: "future smaller delta", Expected: d, Actual: d + time.Millisecond*2, Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "future border delta", Expected: d, Actual: d + time.Millisecond*3, Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "future larger delta", Expected: d, Actual: d + time.Millisecond*4, Pass: false, Epsilon: time.Millisecond * 3},
		// Test cases for the actual time being in the past
		{Name: "past smaller delta", Expected: d, Actual: d - time.Millisecond*2, Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "past border delta", Expected: d, Actual: d - time.Millisecond*3, Pass: true, Epsilon: time.Millisecond * 3},
		{Name: "past larger delta", Expected: d, Actual: d - time.Millisecond*4, Pass: false, Epsilon: time.Millisecond * 3},
	}

	for i, tc := range c {
		pass := ExpectedApprox(&tc, tc.Expected, tc.Actual, tc.Epsilon, tc.Name)
		if pass != tc.Pass {
			t.Errorf(`[%d] %s failed: "%s"`, i, tc.Name, tc.loggedMessage)
		}
	}
}

func TestExpectedApprox(t *testing.T) {
	f1 := 0.15
	c := []testCaseApprox[float64, float64]{
		// Exactness, with 0 epsilon
		{Name: "same value", Expected: f1, Actual: 0.15, Pass: true},
		// Test cases for the actual time being in the future
		{Name: "add smaller delta", Expected: f1, Actual: f1 + 0.2, Pass: true, Epsilon: 0.3},
		{Name: "add border delta", Expected: f1, Actual: f1 + 0.3, Pass: true, Epsilon: 0.3},
		{Name: "add larger delta", Expected: f1, Actual: f1 + 0.4, Pass: false, Epsilon: 0.3},
		// Test cases for the actual time being in the past
		{Name: "sub smaller delta", Expected: f1, Actual: f1 - 0.2, Pass: true, Epsilon: 0.3},
		{Name: "sub border delta", Expected: f1, Actual: f1 - 0.3, Pass: true, Epsilon: 0.3},
		{Name: "sub larger delta", Expected: f1, Actual: f1 - 0.4, Pass: false, Epsilon: 0.3},
	}

	for i, tc := range c {
		pass := ExpectedApprox(&tc, tc.Expected, tc.Actual, tc.Epsilon, tc.Name)
		if pass != tc.Pass {
			t.Errorf(`[%d] %s failed: "%s"`, i, tc.Name, tc.loggedMessage)
		}
	}
}
