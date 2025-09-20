package medium

// https://leetcode.com/problems/triangle
func minimumTotal(triangle [][]int) int {
	dp := triangle[len(triangle)-1]

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}
