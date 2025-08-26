package easy

// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle
func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiag := 0
	maxArea := 0

	for _, d := range dimensions {
		// computing squared diagonal is enough
		diag := d[0]*d[0] + d[1]*d[1]
		area := d[0] * d[1]
		if diag > maxDiag || (diag == maxDiag && area > maxArea) {
			maxDiag = diag
			maxArea = area
		}
	}

	return maxArea
}
