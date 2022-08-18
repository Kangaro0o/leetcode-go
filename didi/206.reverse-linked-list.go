package didi

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1}
	dummyHead.Next = head
	p := head
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
	node := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}
