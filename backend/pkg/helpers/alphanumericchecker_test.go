package helpers_test

import (
	"socialnetwork/pkg/helpers"
	"testing"
)

func TestIsAlphaNumeric(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"test123", true},
		{"TEST123", true},
		{"Test123", true},
		{"123", true},
		{"abc", true},
		{"test 123", false}, // contains space
		{"test!123", false}, // contains special character
		{"", true},          // empty string
		{" ", false},        // only space
		{"!@#^&*", false},   // special characters
		{"Test 123", false}, // contains space
		{"Test!123", false}, // contains special character
		{"123 456", false},  // contains space
	}

	for _, c := range cases {
		got, err := helpers.IsAlphaNumeric(c.in)
		if err != nil {
			t.Error(err)
		}
		if got != c.want {
			t.Errorf("isAlphanumeric(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
