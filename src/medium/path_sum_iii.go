package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSumTraverse(root *TreeNode, targetSum int, accumulated []int) int {
	if root == nil {
		return 0
	}

	total := 0

	s := root.Val
	if s == targetSum {
		total++
	}

	for i := len(accumulated) - 1; i >= 0; i-- {
		s += accumulated[i]
		if s == targetSum {
			total++
		}
	}

	accumulated = append(accumulated, root.Val)

	return total +
		pathSumTraverse(root.Left, targetSum, accumulated) +
		pathSumTraverse(root.Right, targetSum, accumulated)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		if root.Val == targetSum {
			return 1
		}
		return 0
	}

	return pathSumTraverse(root, targetSum, nil)
}
