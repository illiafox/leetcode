package hard

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVerticalTraversal(t *testing.T) {
	tests := []struct {
		input    []any
		expected [][]int
	}{
		{
			input:    []any{3, 9, 20, nil, nil, 15, 7},
			expected: [][]int{{9}, {3, 15}, {20}, {7}},
		},
		{
			input:    []any{},
			expected: nil,
		},
		{
			input:    []any{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
		{
			input:    []any{1, 2, 3, 4, 6, 5, 7},
			expected: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}},
		},
	}

	// all implementations
	funcs := []func(*TreeNode) [][]int{verticalTraversal, verticalTraversalFirstTry}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			for _, f := range funcs {
				root := BuildTree(tt.input)
				result := f(root)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("verticalTraversal(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}
