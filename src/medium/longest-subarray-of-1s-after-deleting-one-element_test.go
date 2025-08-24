package medium

import (
	"fmt"
	"testing"
)

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		// Examples from the prompt
		{nums: []int{1, 1, 0, 1}, expected: 3},
		{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, expected: 5},
		{nums: []int{1, 1, 1}, expected: 2},

		// Edge cases
		{nums: []int{}, expected: 0},           // empty
		{nums: []int{1}, expected: 0},          // single 1 (must delete it)
		{nums: []int{0}, expected: 0},          // single 0
		{nums: []int{0, 0, 0}, expected: 0},    // all zeros
		{nums: []int{1, 1, 1, 1}, expected: 3}, // all ones -> delete one

		// Zeros at boundaries
		{nums: []int{0, 1, 1, 1, 1, 1, 0}, expected: 5},
		{nums: []int{0, 0, 1, 1}, expected: 2},
		{nums: []int{1, 1, 0, 0, 1, 1, 1}, expected: 3},

		// Alternating and multiple zeros
		{nums: []int{1, 0, 1, 0, 1}, expected: 2},
		{nums: []int{1, 0, 1, 1, 1}, expected: 4},             // delete the single zero to merge
		{nums: []int{1, 0, 0, 1, 1, 1, 0, 1, 1}, expected: 5}, // delete the zero between runs
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.nums), func(t *testing.T) {
			got := longestSubarray(tc.nums)
			if got != tc.expected {
				t.Fatalf("longestSubarray(%v)=%d, want %d", tc.nums, got, tc.expected)
			}
		})
	}
}
