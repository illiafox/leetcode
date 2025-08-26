package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaOfMaxDiagonal(t *testing.T) {
	cases := []struct {
		dimensions [][]int
		expected   int
	}{
		// Single rectangle
		{[][]int{{3, 4}}, 12}, // diag^2=25, area=12

		// Two rectangles, different diagonals
		{[][]int{{3, 4}, {6, 8}}, 48}, // 6x8 has larger diag (100 vs 25)

		// Equal diagonals, different areas
		{[][]int{{5, 12}, {9, 10}}, 90}, // both diag^2=169, area 60 vs 90

		// Multiple rectangles, pick max diag
		{[][]int{{2, 2}, {10, 1}, {6, 8}}, 10}, // 6x8 diag^2=100 is largest

		// Equal diagonals and equal area
		{[][]int{{4, 3}, {3, 4}}, 12}, // both diag^2=25, area=12
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, areaOfMaxDiagonal(tc.dimensions),
			"input = %v", tc.dimensions,
		)
	}
}
