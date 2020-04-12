package model

import (
	"testing"
)

func TestGetZoneFromHost(t *testing.T) {

}
func TestGetReversedDomain(t *testing.T) {
	// ensures that domain names are properly converted to slices and reversed
	table := []struct {
		in  string
		out []string
	}{
		{"test.bor", []string{"bor", "test"}},
		{"longer.test.bor", []string{"bor", "test", "longer"}},
	}

	for _, x := range table {
		result := getReversedDomain(x.in)

		if len(result) != len(x.out) {
			t.Errorf("expected: %q for %q", x.out, result)
		}

		for i := 0; i < len(x.out); i++ {
			if result[i] != x.out[i] {
				t.Errorf("expected: %q for %q", x.out, result)
			}
		}

	}

}
