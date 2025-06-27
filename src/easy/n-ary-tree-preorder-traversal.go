package easy

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	s := make([]int, 0, len(root.Children))

	var process func(node *Node)
	process = func(node *Node) {
		if node == nil {
			return
		}

		s = append(s, node.Val)

		for _, n := range node.Children {
			process(n)
		}
	}

	process(root)

	return s
}
