package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/find-mode-in-binary-search-tree/
func findMode(root *TreeNode) []int {
	freq := make(map[int]int)

	var process func(node *TreeNode)
	process = func(node *TreeNode) {
		if node == nil {
			return
		}

		freq[node.Val]++
		process(node.Left)
		process(node.Right)
	}

	process(root)

	maxFreq := 0
	var modes []int

	for val, count := range freq {
		if count > maxFreq {
			maxFreq = count
			modes = append(modes[:0], val)
		} else if count == maxFreq {
			modes = append(modes, val)
		}
	}

	return modes
}
