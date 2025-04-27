package easy

import "math"

func getMinimumDifferenceTraverse(node *TreeNode, prev *int, minDiff *int) {
	if node == nil {
		return
	}
	getMinimumDifferenceTraverse(node.Left, prev, minDiff)
	if *prev != -1 && node.Val-*prev < *minDiff {
		*minDiff = node.Val - *prev
	}
	*prev = node.Val
	getMinimumDifferenceTraverse(node.Right, prev, minDiff)
}

func getMinimumDifference(node *TreeNode) int {
	minDiff := math.MaxInt64
	prev := -1

	getMinimumDifferenceTraverse(node, &prev, &minDiff)
	return minDiff
}
