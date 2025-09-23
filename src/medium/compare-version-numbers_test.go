package medium

import (
	"testing"
)

func TestCompareVersion(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"1.0", "1", 0},
		{"1.0.0", "1.0", 0},
		{"1.01", "1.001", 0},
		{"1.0.1", "1", 1},
		{"1.0.0.0.2", "1.0.0.0.1", 1},
		{"2.0", "2.0.1", -1},
		{"0.0.0", "0", 0},
		{"01.002", "1.2", 0}, // leading zeros
	}

	for _, tc := range cases {
		got := compareVersion(tc.a, tc.b)
		if got != tc.want {
			t.Fatalf("compareVersion=%d, want %d; a=%v; b=%v", got, tc.want, tc.a, tc.b)
		}
	}
}
