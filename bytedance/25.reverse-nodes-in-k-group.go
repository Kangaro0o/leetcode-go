package bytedance

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Val: -1}
	dummyHead.Next = head
	return nil
}
