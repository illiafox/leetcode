package medium

// https://leetcode.com/problems/unique-binary-search-trees
// https://en.wikipedia.org/wiki/Catalan_number
// https://www.geeksforgeeks.org/applications-of-catalan-numbers/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}
