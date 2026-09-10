package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		curr := queue[0]
		queue = queue[1:]

		if i < len(values) {
			if values[i] != nil {
				curr.Left = &TreeNode{Val: *values[i]}
				queue = append(queue, curr.Left)
			}
			i++
		}

		if i < len(values) {
			if values[i] != nil {
				curr.Right = &TreeNode{Val: *values[i]}
				queue = append(queue, curr.Right)
			}
			i++
		}
	}

	return root
}
