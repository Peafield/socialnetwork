package helpers_test

import (
	"socialnetwork/pkg/helpers"
	"testing"
)

func TestIsValidPath(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"../helpers", true},
		{"./helpers01", false},
		{"../db/migrations", true},
		{"../models/migrations", false},
	}
	for _, c := range cases {
		got, err := helpers.IsValidPath(c.in)
		if err != nil {
			t.Log(err)
		}
		if got != c.want {
			t.Errorf("IsValidPath(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
