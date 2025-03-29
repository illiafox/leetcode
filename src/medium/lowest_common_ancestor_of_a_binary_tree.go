package medium

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root, target *TreeNode, path []*TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root == target {
		return path
	}

	path = append(path, root)

	if res := traverse(root.Left, target, path); res != nil {
		return res
	}

	return traverse(root.Right, target, path)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := traverse(root, p, nil)
	pathQ := traverse(root, q, nil)

	if slices.Contains(pathQ, p) {
		return p
	}
	if slices.Contains(pathP, q) {
		return q
	}

	for i := len(pathP) - 1; i >= 0; i-- {
		ptr := pathP[i]

		if slices.Contains(pathQ, ptr) {
			return ptr
		}
	}

	return root
}
