package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.
// https://leetcode.com/problems/partition-list
func partition(head *ListNode, x int) *ListNode {
	lessDummy := new(ListNode)
	moreDummy := new(ListNode)

	lessTail, moreTail := lessDummy, moreDummy
	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			lessTail.Next = node
			lessTail = lessTail.Next
		} else {
			moreTail.Next = node
			moreTail = moreTail.Next
		}
	}

	moreTail.Next = nil
	lessTail.Next = moreDummy.Next

	return lessDummy.Next
}
