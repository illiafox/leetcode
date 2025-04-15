package medium

type BackList struct {
	Prev *BackList
	Val  int
}

func getBackList(prev *BackList, length int) []int {
	out := make([]int, length)

	for node := prev; node != nil; node = node.Prev {
		length--
		out[length] = node.Val
	}

	return out
}

func pathSum2Traverse(root *TreeNode, accum *[][]int, curSum, targetSum int, prev *BackList, length int) {
	if root == nil {
		return
	}

	curSum += root.Val

	prev = &BackList{
		Prev: prev,
		Val:  root.Val,
	}

	if root.Right == nil && root.Left == nil {
		if curSum == targetSum {
			*accum = append(*accum, getBackList(prev, length))
		}
		return
	}

	length++

	pathSum2Traverse(root.Left, accum, curSum, targetSum, prev, length)
	pathSum2Traverse(root.Right, accum, curSum, targetSum, prev, length)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func path2Sum(root *TreeNode, targetSum int) [][]int {
	out := make([][]int, 0)

	pathSum2Traverse(root, &out, 0, targetSum, nil, 1)
	return out
}
