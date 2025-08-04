package hard

/*
// https://leetcode.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	indices := make([]int, len(heights)+2)
	copy(indices[1:], heights)

	stack := []int{0}
	maxArea := 0

	for i := 0; i < len(indices); i++ {
		for len(stack) > 0 && indices[stack[len(stack)-1]] > indices[i] {
			h := indices[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}

			area := h * w
			maxArea = max(maxArea, area)
		}
		stack = append(stack, i)
	}

	return maxArea
}
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))

	maxArea := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}

	return maxArea
}
