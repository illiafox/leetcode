package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpGreedy(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}

	cases := []testCase{
		{[]int{0}, 0},             // single element
		{[]int{1, 0}, 1},          // one jump to end
		{[]int{2, 3, 1, 1, 4}, 2}, // classic example
		{[]int{1, 1, 1, 1}, 3},    // must take each step
		{[]int{4, 1, 1, 1, 1}, 1}, // big first jump reaches end
		{[]int{2, 0, 2, 0, 1}, 2}, // zeros but still reachable
		{[]int{1, 2, 1, 1, 1}, 3}, // choose farther reach from layer
		{[]int{3, 2, 1}, 1},       // first jump already past end
		{[]int{2, 1}, 1},          // minimal non-trivial
		{[]int{1, 2, 3}, 2},       // greedy expands layers correctly
	}

	for _, c := range cases {
		got := jump(c.nums)
		assert.Equal(t, c.expected, got, "nums=%v", c.nums)
	}
}
