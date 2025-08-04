package hard

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{1, 1}, 2},
		{[]int{6, 2, 5, 4, 5, 1, 6}, 12},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{0, 9}, 9},
		{[]int{2, 1, 2}, 3},
		{[]int{4, 2, 0, 3, 2, 5}, 6},
	}

	for _, tt := range tests {
		got := largestRectangleArea(tt.heights)
		if got != tt.want {
			t.Errorf("largestRectangleArea(%v) = %d; want %d", tt.heights, got, tt.want)
		}
	}
}
