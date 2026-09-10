package medium

import (
	"slices"
	"testing"
)

func listFromSlice(values []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for _, val := range values {
		node := &ListNode{Val: val}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		l1, l2   []int
		expected []int
	}{
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			l1:       []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2:       []int{5, 6, 4},
			expected: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, c := range cases {
		res := addTwoNumbers(listFromSlice(c.l1), listFromSlice(c.l2))

		if !slices.Equal(listToSlice(res), c.expected) {
			t.Error("case failed: expected", c.expected, "got", listToSlice(res))
		}
	}
}
