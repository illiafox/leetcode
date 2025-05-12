package medium

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// solution for https://leetcode.com/problems/find-duplicate-subtrees/solutions/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var out []*TreeNode

	count := map[string]int{}

	var traverse func(*TreeNode) string
	traverse = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		left := traverse(node.Left)
		right := traverse(node.Right)
		key := fmt.Sprintf("%d,%s,%s", node.Val, left, right)

		count[key]++
		if count[key] == 2 {
			out = append(out, node)
		}
		return key
	}

	traverse(root)
	return out
}
