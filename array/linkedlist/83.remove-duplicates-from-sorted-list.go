package linkedlist

// deleteDuplicates 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	curr := head
	for head.Next != nil { // 遍历链表
		next := head.Next
		if curr.Val == next.Val {
			second := next.Next
			curr.Next = second // 删除 next 节点
		} else {
			curr = next
		}
		head = curr
	}
	return dummy.Next
}
