package easy

import "math/bits"

// https://leetcode.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
