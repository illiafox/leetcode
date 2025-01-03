package medium

import (
	"reflect"
	"testing"
)

func createListFromArray(values []int) *ListNode {
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

func listToArray(head *ListNode) []int {
	var result []int

	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}

	return result
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
		res := addTwoNumbers(createListFromArray(c.l1), createListFromArray(c.l2))

		if !reflect.DeepEqual(listToArray(res), c.expected) {
			t.Error("case failed: expected", c.expected, "got", listToArray(res))
		}
	}
}
