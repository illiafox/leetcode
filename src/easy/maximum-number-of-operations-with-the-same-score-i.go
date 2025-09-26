package easy

// https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-i/
func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sum := nums[0] + nums[1]
	count := 1

	for i := 2; i < len(nums)-1; i += 2 {
		if sum != nums[i]+nums[i+1] {
			break
		}
		count++
	}
	return count
}
