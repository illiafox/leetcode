package hard

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCost(t *testing.T) {
	type testCase struct {
		grid     [][]int
		expected int
	}

	var cases = []testCase{
		{
			grid:     [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}},
			expected: 3,
		},
		{
			grid:     [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}},
			expected: 0,
		},
		{
			grid:     [][]int{{1, 2}, {4, 3}},
			expected: 1,
		},
		{
			grid:     [][]int{{2, 2, 2}, {2, 2, 2}},
			expected: 3,
		},
		{
			grid: [][]int{
				{3, 4, 3}, {2, 2, 2}, {2, 1, 1}, {4, 3, 2}, {2, 1, 4}, {2, 4, 1},
				{3, 3, 3}, {1, 4, 2}, {2, 2, 1}, {2, 1, 1}, {3, 3, 1}, {4, 1, 4},
				{2, 1, 4}, {3, 2, 2}, {3, 3, 1}, {4, 4, 1}, {1, 2, 2}, {1, 1, 1},
				{1, 3, 4}, {1, 2, 1}, {2, 2, 4}, {2, 1, 3}, {1, 2, 1}, {4, 3, 2},
				{3, 3, 4}, {2, 2, 1}, {3, 4, 3}, {4, 2, 3}, {4, 4, 4},
			},
			expected: 18,
		},
	}

	for _, c := range cases {
		assert.Equal(t, minCost(c.grid), c.expected, fmt.Sprintf("minCost(%v)", c.grid))
	}
}
