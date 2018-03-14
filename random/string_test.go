package random

import (
	"testing"
)

func TestNewString(t *testing.T) {
	tests := []struct {
		length int
		set    string
	}{
		{8, "0123456789"},
		{10, "abcdefghijklmnopqrstuvwxyz"},
		{27, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{37, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"},
		{55, "Hello, world!"},
	}
	for _, test := range tests {
		r := NewString(test.length, test.set)
		t.Log(r)
		if (test.length != len(r)) {
			t.Error("length is ", test.length, "len(r) is ", len(r))
		}
	}
}
func TestNewString2(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(NewString(13, `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789,.?/!@#$%^&*()_+-=<>{}[]|\'"`))
	}
}

func TestNewCryptoString(t *testing.T) {
	for i := 0; i < 5; i++ {
		r, err := NewCryptoString(i + 1)
		if err != nil {
			t.Log(err)
		} else {
			t.Log("length is", len(r), ", result is", r)
		}
	}
}

func TestNewUUID(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(NewUUID())
	}
}
