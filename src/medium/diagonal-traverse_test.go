package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindDiagonalOrder_Table(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want []int
	}{
		{
			name: "single cell",
			mat:  [][]int{{42}},
			want: []int{42},
		},
		{
			name: "1x3 row",
			mat:  [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "3x1 column",
			mat:  [][]int{{1}, {2}, {3}},
			want: []int{1, 2, 3},
		},
		{
			name: "2x2 square",
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 3, 4}, // diagonals: [1], [2,3], [4] with zig-zag
		},
		{
			name: "3x3 square",
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "2x3 rectangle",
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: []int{1, 2, 4, 5, 3, 6},
		},
		{
			name: "3x2 rectangle",
			mat: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			want: []int{1, 2, 3, 5, 4, 6},
		},
		{
			name: "1x4 row",
			mat:  [][]int{{-1, 0, 7, 8}},
			want: []int{-1, 0, 7, 8},
		},
		{
			name: "4x1 column",
			mat:  [][]int{{-1}, {0}, {7}, {8}},
			want: []int{-1, 0, 7, 8},
		},
		{
			name: "mixed values 3x4",
			mat: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
			},
			want: []int{5, 1, 2, 13, 4, 9, 11, 8, 3, 6, 10, 7},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findDiagonalOrder(tc.mat)
			require.Equal(t, tc.want, got, "mat=%v", tc.mat)
		})
	}
}
