package linkedlist

// reverseBetween 92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	p := dummy
	n := left - 1
	for n > 0 {
		p = p.Next
		n--
	}
	//t := p
	headNode, _ := reverseKList(p.Next, right-left)
	p.Next = headNode.Next
	return dummy.Next
}

func reverseKList(head *ListNode, k int) (*ListNode, *ListNode) {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	prev := dummy
	curr := head
	for i := 0; i < k; i++ {
		next := curr.Next
		second := next.Next
		next.Next = prev.Next
		curr.Next = second
		prev.Next = next // prev 重新指向 子链表 头节点
	}
	return prev, curr
}
