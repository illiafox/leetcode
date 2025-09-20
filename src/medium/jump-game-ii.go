package medium

// https://leetcode.com/problems/jump-game-ii
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	steps, curEnd, far := 0, 0, 0
	for i := 0; i < n; i++ {
		if distance := i + nums[i]; distance > far {
			far = distance
		}
		if i == curEnd {
			steps++
			curEnd = far
			if curEnd >= n-1 {
				break
			}
		}
	}

	return steps
}
