package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacificAtlantic(t *testing.T) {
	type testCase struct {
		heights [][]int
		want    [][]int
	}

	cases := []testCase{
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		{
			heights: [][]int{
				{1, 1},
				{1, 1},
			},
			want: [][]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
			},
		},
		{
			heights: [][]int{
				{3, 3, 3},
				{3, 1, 3},
			},
			want: [][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{1, 0},
				{1, 2},
			},
		},
	}

	for _, tc := range cases {
		got := pacificAtlantic(tc.heights)
		assert.ElementsMatch(t, tc.want, got)
	}
}
