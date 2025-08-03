package hard

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"-(3+2)", -5},
		{"1-(2+3)", -4},
		{"1-(-2)", 3},
		{"10 + (2 - (3 + (4 - 5)))", 10},
		{"(1)", 1},
		{"-1 + (2 - 3)", -2},
		{"(7)-(0)+(4)", 11},
		{"(1+(4+5+2)-3)+(6+8) + 100", 123},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := calculate(tt.input)
			if result != tt.expected {
				t.Errorf("calculate(%q) = %d; want %d", tt.input, result, tt.expected)
			}

			result = calculateOld(tt.input)
			if result != tt.expected {
				t.Errorf("calculateOld(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
