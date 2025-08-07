package hard

import (
	"testing"
)

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestBinaryTreeEncoder(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
	}{
		{
			name: "Empty tree",
			tree: nil,
		},
		{
			name: "Single node",
			tree: &TreeNode{Val: 42},
		},
		{
			name: "Complete binary tree",
			tree: BuildTree([]any{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			name: "Left-skewed tree",
			tree: BuildTree([]any{1, 2, nil, 3, nil, 4}),
		},
		{
			name: "Right-skewed tree",
			tree: BuildTree([]any{1, nil, 2, nil, 3, nil, 4}),
		},
		{
			name: "Sparse tree",
			tree: BuildTree([]any{1, nil, 2, 3, nil, nil, 4}),
		},
	}

	encoder := NewCodec()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialized := encoder.serialize(tt.tree)
			deserialized := encoder.deserialize(serialized)

			if !isSameTree(tt.tree, deserialized) {
				t.Errorf("Mismatch in tree after serialize/deserialize.\nOriginal: %+v\nResult: %+v\nSerialized: %s",
					tt.tree, deserialized, serialized)
			}
		})
	}
}
