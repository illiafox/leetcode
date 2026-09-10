package medium

import (
	"slices"
	"testing"
)

func listToSlice(head *ListNode) []int {
	var result []int
	for curr := head; curr != nil; curr = curr.Next {
		result = append(result, curr.Val)
	}
	return result
}

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "single node",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two nodes (even)",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "three nodes (odd)",
			input:    []int{1, 2, 3},
			expected: []int{2, 1, 3},
		},
		{
			name:     "four nodes (even)",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			name:     "five nodes (odd)",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 1, 4, 3, 5},
		},
	}

	funcs := []struct {
		name string
		f    func(head *ListNode) *ListNode
	}{
		{name: "swapPairs/", f: swapPairs},
		{name: "swapPairsSimplified/", f: swapPairsSimplified},
	}

	for _, f := range funcs {
		for _, tc := range tests {
			t.Run(f.name+tc.name, func(t *testing.T) {
				list := listFromSlice(tc.input)
				swapped := swapPairs(list)
				actual := listToSlice(swapped)

				if !slices.Equal(actual, tc.expected) {
					t.Errorf("swapPairs(%v) = %v, want %v", tc.input, actual, tc.expected)
				}
			})
		}
	}
}
