package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/maximum-binary-tree
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	return &TreeNode{
		Val:   nums[maxIdx],
		Left:  constructMaximumBinaryTree(nums[:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
}
