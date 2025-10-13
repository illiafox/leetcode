package easy

// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i
func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if k <= 0 || 2*k > n {
		return false
	}

	increment := make([]int, n)
	increment[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			increment[i] = increment[i+1] + 1
		} else {
			increment[i] = 1
		}
	}

	for i := 0; i <= n-2*k; i++ {
		if increment[i] >= k && increment[i+k] >= k {
			return true
		}
	}

	return false
}
