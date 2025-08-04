package hard

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
