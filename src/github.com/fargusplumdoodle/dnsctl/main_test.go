package main

import (
	"testing"
)

func TestGetURLWithTrailingSlash(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{"http://localhost/", "http://localhost/"},
		{"http://localhost", "http://localhost/"},
	}

	for _, x := range table {
		result := GetURLWithTrailingSlash(x.in)

		if result != x.out {
			t.Errorf("expected: %q got %q", x.out, result)
		}
	}
}
