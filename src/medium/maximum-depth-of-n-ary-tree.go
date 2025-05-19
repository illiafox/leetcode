package medium

type Node struct {
	Val      int
	Children []*Node
}

func maxDepthTraverse(out *int, node *Node, depth int) {
	if node == nil {
		return
	}

	depth++

	if len(node.Children) == 0 {
		*out = max(*out, depth)
		return
	}

	for _, n := range node.Children {
		maxDepthTraverse(out, n, depth)
	}
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func maxDepth(root *Node) int {
	m := 0
	maxDepthTraverse(&m, root, 0)

	return m
}
