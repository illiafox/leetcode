package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "LeetCode Example 1: Standard repeating mix",
			input:    "abcabc",
			expected: 10,
		},
		{
			name:     "LeetCode Example 2: Prefixed duplicates",
			input:    "aaacb",
			expected: 3,
		},
		{
			name:     "LeetCode Example 3: Minimum valid string",
			input:    "abc",
			expected: 1,
		},
		{
			name:     "Edge Case: Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Edge Case: Missing one required character",
			input:    "aaaaaaabbbbbbb",
			expected: 0,
		},
		{
			name:     "Edge Case: Short string with missing characters",
			input:    "ab",
			expected: 0,
		},
		{
			name:     "Longer valid string",
			input:    "abccba",
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := numberOfSubstrings(tt.input)

			assert.Equal(t, tt.expected, actual,
				"Input %q, expected %d, got %d",
				tt.input, tt.expected, actual,
			)
		})
	}
}
