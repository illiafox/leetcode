package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/house-robber-iii/
func robTraverse(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	lRob, lSkip := robTraverse(node.Left)
	rRob, rSkip := robTraverse(node.Right)

	rob := node.Val + lSkip + rSkip
	skip := max(lRob, lSkip) + max(rRob, rSkip)

	return rob, skip
}

func rob(root *TreeNode) int {
	return max(robTraverse(root))
}
