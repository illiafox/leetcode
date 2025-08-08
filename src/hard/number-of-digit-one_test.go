package hard

import "testing"

func TestCountDigitOne(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{5, 1},
		{10, 2}, // 1, 10
		{13, 6}, // 1, 10, 11, 12, 13
		{20, 12},
		{99, 20},
		{100, 21},
		{101, 23},
		{110, 33},
		{123, 57},
		{1234, 689},
		{9999, 4000},
		{10000, 4001},
		{100000, 50001},
	}

	for _, tt := range tests {
		got := countDigitOne(tt.n)
		if got != tt.expected {
			t.Errorf("countDigitOne(%d) = %d; want %d", tt.n, got, tt.expected)
		}
	}
}
