package medium

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	minLength := math.MaxInt

	start := 0
	sum := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			minLength = min(minLength, end-start+1)
			sum -= nums[start]
			start++
		}
	}
	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}
