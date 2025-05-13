package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var out []float64

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0

		levelLength := len(queue)
		for range levelLength {
			pop := queue[0]
			queue = queue[1:] // TODO: use deque

			sum += pop.Val
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
		}
		out = append(out, float64(sum)/float64(levelLength))
	}

	return out
}
