package signing

import (
	"testing"
)

const (
	dataToBeSigned   = "My Account Info Here"
	expectedHashSize = 32
)

func TestIsValid(t *testing.T) {
	key, err := NewHMACKey()
	if err != nil {
		t.FailNow()
	}

	// Reset the timer, since we don't want to time the setup we had to do
	sig := key.Sign(dataToBeSigned)
	if len(sig) != expectedHashSize {
		t.Error("Signature not expected length")
	}

	verified := key.IsValid(dataToBeSigned, sig)
	if !verified {
		t.Error("Failed to verify expected matching signatures")
	}

	verified = key.IsValid(dataToBeSigned+"1", sig)
	if verified {
		t.Error("Improperly verified differing signatures")
	}
}

func BenchmarkSignHMAC(b *testing.B) {
	key, err := NewHMACKey()
	if err != nil {
		b.FailNow()
	}

	// Reset the timer, since we don't want to time the setup we had to do
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key.Sign(dataToBeSigned)
	}
}

func BenchmarkVerifyHMAC(b *testing.B) {
	key, err := NewHMACKey()
	if err != nil {
		b.FailNow()
	}
	sig := key.Sign(dataToBeSigned)

	// Reset the timer, since we don't want to time the setup we had to do
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		key.IsValid(dataToBeSigned, sig)
	}
}
