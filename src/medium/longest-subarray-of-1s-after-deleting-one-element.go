package medium

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element
func longestSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	maxLength := 0

	length := 0
	zeroFound := false
	afterZero := 0

	for _, num := range nums {
		switch num {
		case 1:
			length++
			if zeroFound {
				afterZero++
			}
		case 0:
			if !zeroFound {
				zeroFound = true
				continue
			}

			maxLength = max(length, maxLength)
			length = afterZero
			afterZero = 0
		}
	}

	if !zeroFound {
		// one element must be deleted
		return len(nums) - 1
	}

	return max(length, maxLength)
}
