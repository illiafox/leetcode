package hard

import (
	"encoding/json"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
type BinaryTreeEncoder struct{}

func NewCodec() BinaryTreeEncoder {
	return BinaryTreeEncoder{}
}

// Serializes a tree to a single string.
func (c *BinaryTreeEncoder) serialize(root *TreeNode) string {
	if root == nil {
		return "[null]"
	}

	vals := []*int{&root.Val}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				vals = append(vals, &node.Left.Val)
				queue = append(queue, node.Left)
			} else {
				vals = append(vals, nil)
			}
			if node.Right != nil {
				vals = append(vals, &node.Right.Val)
				queue = append(queue, node.Right)
			} else {
				vals = append(vals, nil)
			}
		}

		queue = queue[l:]
	}

	// trim nils
	for len(vals) > 0 && vals[len(vals)-1] == nil {
		vals = vals[:len(vals)-1]
	}

	data, err := json.Marshal(vals)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// Deserializes your encoded data to tree.
func (c *BinaryTreeEncoder) deserialize(data string) *TreeNode {
	var vals []*int
	if err := json.Unmarshal([]byte(data), &vals); err != nil {
		panic(err)
	}

	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}

	i := 1
	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// left child
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Left)
			}
			i++
		}

		// right child
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
