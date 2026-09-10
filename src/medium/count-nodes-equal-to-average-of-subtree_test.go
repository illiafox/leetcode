package medium

import (
	"testing"
)

func TestAverageOfSubtree(t *testing.T) {
	tests := []struct {
		name     string
		input    []*int
		expected int
	}{
		{
			name:     "leetcode example 1",
			input:    []*int{new(4), new(8), new(5), new(0), new(1), nil, new(6)},
			expected: 5,
		},
		{
			name:     "leetcode example 2",
			input:    []*int{new(1)},
			expected: 1,
		},
		{
			name:     "empty tree",
			input:    []*int{},
			expected: 0,
		},
		{
			name:     "unbalanced",
			input:    []*int{new(10), new(5), nil, new(2)},
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := BuildTree(tc.input)
			actual := averageOfSubtree(root)
			if actual != tc.expected {
				t.Fatalf("averageOfSubtree() = %d, expected %d", actual, tc.expected)
			}
		})
	}
}
