package easy

import "math"

// https://leetcode.com/problems/valid-perfect-square
func isPerfectSquare(num int) bool {
	start, end := 1, num

	for start <= end {
		mid := start + (end-start)/2
		if mid*mid == num {
			return true
		}

		if mid*mid < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}

func isPerfectSquareMath(num int) bool {
	s := math.Sqrt(float64(num))
	return s == math.Trunc(s)
}
