package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// For a valid binary-search tree, an in-order traversal ( left → visit → right ) outputs the keys in strictly ascending order.
// Therefore, when two keys are swapped, the in-order sequence will contain one or two inversions—places where a value is smaller than the one that precedes it.
func recoverTreeTraverse(node *TreeNode, prev **TreeNode, badNodes *[]*TreeNode) {
	if node == nil {
		return
	}
	recoverTreeTraverse(node.Left, prev, badNodes)

	if *prev != nil && (*prev).Val > node.Val {
		*badNodes = append(*badNodes, *prev)
		*badNodes = append(*badNodes, node)
	}
	*prev = node

	recoverTreeTraverse(node.Right, prev, badNodes)
}

// https://leetcode.com/problems/recover-binary-search-tree/
func recoverTree(root *TreeNode) {
	var badNodes []*TreeNode
	var prev *TreeNode
	recoverTreeTraverse(root, &prev, &badNodes)

	if len(badNodes) < 2 {
		panic("bst does not have enough bad nodes swap")
	}

	badNodes[0].Val, badNodes[len(badNodes)-1].Val = badNodes[len(badNodes)-1].Val, badNodes[0].Val
}
