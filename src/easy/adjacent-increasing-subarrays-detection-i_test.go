package easy

import "testing"

func TestHasIncreasingSubarrays(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			// [1,2,3] and [4,5,6] are both strictly increasing
			want: true,
		},
		{
			nums: []int{10, 0},
			k:    1,
			// two adjacent singletons -> trivially true
			want: true,
		},
		{
			nums: []int{1, 2, 3},
			k:    2,
			// need 4 elements for two adjacent length-2 windows
			want: false,
		},
		{
			nums: []int{1, 2, 3, 4},
			k:    0,
			want: false,
		},
		{
			nums: []int{1, 2, 2, 3, 4, 5},
			k:    2,
			// possible windows starting at 3: [3,4] and [4,5] => true
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4},
			k:    2,
			// [1,2] and [3,4] -> both increasing
			want: true,
		},
		{
			nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
			k:    5,
			want: false,
		},
		{
			nums: []int{-15, 19},
			k:    1,
			want: true,
		},
		{
			nums: []int{5, 8, -2, -1},
			k:    2,
			want: true,
		},
	}

	for _, tc := range cases {
		got := hasIncreasingSubarrays(tc.nums, tc.k)
		if got != tc.want {
			t.Fatalf("hasIncreasingSubarrays(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
		}
	}
}
