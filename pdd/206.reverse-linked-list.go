package pdd

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Val: -1}
	dummyHead.Next = head
	p := dummyHead.Next

	for p.Next != nil {
		next := p.Next
		second := next.Next
		next.Next = dummyHead.Next
		dummyHead.Next = next
		p.Next = second
	}
	return dummyHead.Next
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
