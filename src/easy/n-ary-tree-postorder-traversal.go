package easy

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	s := make([]int, 0, len(root.Children))

	var process func(node *Node)
	process = func(node *Node) {
		if node == nil {
			return
		}

		for _, n := range node.Children {
			process(n)
		}

		s = append(s, node.Val)
	}

	process(root)

	return s
}
