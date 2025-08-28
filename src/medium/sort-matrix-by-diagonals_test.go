package medium

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMatrix(t *testing.T) {
	tests := []struct {
		in   [][]int
		want [][]int
	}{
		{
			in:   [][]int{},
			want: [][]int{},
		},
		{
			in:   [][]int{{42}},
			want: [][]int{{42}},
		},
		{
			in: [][]int{
				{4, 1},
				{3, 2},
			},
			want: [][]int{
				{4, 1},
				{3, 2},
			},
		},
		{
			in: [][]int{
				{1, 7, 3},
				{9, 8, 2},
				{4, 5, 6},
			},
			want: [][]int{
				{8, 2, 3},
				{9, 6, 7},
				{4, 5, 1},
			},
		},
	}

	for _, tc := range tests {
		cp := deepCopy(tc.in)
		got := sortMatrix(cp)

		assert.Equal(t, tc.want, got)
	}
}

func deepCopy(m [][]int) [][]int {
	out := make([][]int, len(m))
	for i := range m {
		out[i] = slices.Clone(m[i])
	}
	return out
}
