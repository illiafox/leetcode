package medium

// https://leetcode.com/problems/single-element-in-a-sorted-array
func singleNonDuplicate(nums []int) int {
	s := 0
	for i := range nums {
		s ^= nums[i]
	}

	return s
}
