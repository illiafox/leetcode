package hard

import "testing"

func TestIsSelfCrossing(t *testing.T) {
	tests := []struct {
		distance []int
		want     bool
	}{
		// basic no-cross
		{[]int{1, 2, 3, 4}, false},

		// basic cross (forms a square)
		{[]int{2, 1, 1, 2}, true},

		// square loop (self-touching at start)
		{[]int{1, 1, 1, 1}, true},

		// crossing after spiral in
		{[]int{1, 1, 2, 1, 1}, true},

		// longer no-cross spiral
		{[]int{3, 3, 4, 2, 2}, false},

		// touch from behind
		{[]int{1, 1, 1, 2, 1, 1}, true},

		// long no-cross straight line
		{[]int{100, 100, 100, 100, 100}, true},

		// overlapping parallel lines
		{[]int{10, 10, 10, 10, 5, 5, 5, 5}, true},

		// edge case: minimal input
		{[]int{1, 1, 1}, false},

		// edge case: tight touch from back
		{[]int{1, 1, 2, 2, 1, 1}, true},
	}

	for _, tt := range tests {
		got := isSelfCrossing(tt.distance)
		if got != tt.want {
			t.Errorf("isSelfCrossing(%v) = %v; want %v", tt.distance, got, tt.want)
		}
	}
}
