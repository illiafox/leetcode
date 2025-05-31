package easy

import "math"

// https://leetcode.com/problems/valid-perfect-square
// TODO: use binary search?
func isPerfectSquare(num int) bool {
	s := math.Sqrt(float64(num))
	return s == math.Trunc(s)
}
