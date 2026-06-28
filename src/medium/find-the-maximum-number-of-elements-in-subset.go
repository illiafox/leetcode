package medium

import (
	"sort"
)

// https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset
// You are given an array of positive integers nums.
// You need to select a subset of nums which satisfies the following condition:
//
//	You can place the selected elements in a 0-indexed array such that
//	it follows the pattern: [x, x2, x4, ..., xk/2, xk, xk/2, ..., x4, x2, x]
//	(Note that k can be any non-negative power of 2).
//	For example, [2, 4, 16, 4, 2] and [3, 9, 3] follow the pattern while [2, 4, 8, 4, 2] does not.
//
// Return the maximum number of elements in a subset that satisfies these conditions.
//
// 2 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^9
func maximumLength(nums []int) int {
	counts := map[int64]int{}
	for _, num := range nums {
		counts[int64(num)]++
	}

	maxLength := counts[1]
	if maxLength%2 == 0 {
		maxLength--
	}
	delete(counts, 1)

	for num := range counts {
		length := 0
		nextNum := num

		for counts[nextNum] > 1 {
			length += 2
			nextNum *= nextNum
		}

		if _, ok := counts[nextNum]; ok {
			maxLength = max(maxLength, length+1)
		} else {
			maxLength = max(maxLength, length-1)
		}
	}

	return maxLength
}

func maximumLengthSuffix(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}

	ones := counts[1]
	if ones%2 == 0 {
		ones--
	}
	maxSubset := max(1, ones)

	suffix := map[int]int{}
	for i := len(nums) - 1; i >= 0; {
		num := nums[i]

		count := counts[num]
		i -= count

		subset := min(2, count)
		if subset > 1 {
			subset += suffix[num*num]
		}

		suffix[num] = subset

		if subset%2 == 1 {
			maxSubset = max(maxSubset, subset)
		} else {
			maxSubset = max(maxSubset, subset-1)
		}
	}

	return maxSubset
}
