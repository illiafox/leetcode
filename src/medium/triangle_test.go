package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	type testCase struct {
		triangle [][]int
		expected int
	}

	cases := []testCase{
		{[][]int{{5}}, 5},
		{[][]int{{2}, {3, 4}}, 5},
		{[][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}, 11},
		{[][]int{
			{-1},
			{2, 3},
			{1, -1, -3},
		}, -1},
		{[][]int{
			{1},
			{1, 2},
			{4, 3, 1},
		}, 4},
		{[][]int{
			{-10},
			{-5, -6},
			{-3, -4, -2},
		}, -20},
	}

	for _, c := range cases {
		got := minimumTotal(c.triangle)
		assert.Equal(t, c.expected, got)
	}
}
