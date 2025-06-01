package medium

// https://leetcode.com/problems/single-element-in-a-sorted-array
func singleNonDuplicate(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s ^= nums[i]
	}

	return s
}
