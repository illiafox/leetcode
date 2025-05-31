package easy

import "math"

func isPerfectSquare(num int) bool {
	s := math.Sqrt(float64(num))
	return s == math.Trunc(s)
}
