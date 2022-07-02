package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// removeNthFromEnd 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead := head
	length := 0
	for dummyHead != nil {
		dummyHead = dummyHead.Next
		length++
	}
	if length == 1 {
		return nil
	}
	dummyHead = head
	for i := 0; i < length-n-1; i++ {
		dummyHead = dummyHead.Next
	}
	// 此时 dummyHead.Next 就是要删除的元素
	next := dummyHead.Next // 要删除的元素
	if next == nil {
		dummyHead = nil
	} else {
		nn := next.Next
		dummyHead.Next = nn
	}
	return head
}

// 使用快慢指针完成本题（不含虚拟头节点）
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	p1 := head
	// 快指针先走 n 步
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	if p1 == nil { // 如果 p1 走了 n 步后为 nil，则说明 length(list) == n，此时删除头节点即可
		return head.Next
	}
	// 此时慢指针从头开始走，走(length - n)步
	p2 := head
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// 此时 p2 位于要删除元素的上一个节点
	p2.Next = p2.Next.Next
	return head
}

// 使用快慢指针（含虚拟头节点）
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	// 找到倒数第 n+1个节点
	x := findFromEnd(dummy, n+1)
	// 删除下一个节点
	x.Next = x.Next.Next
	return dummy.Next
}

// findFromEnd 查找倒数第 k 个节点
func findFromEnd(dummy *ListNode, k int) *ListNode {
	p1 := dummy
	for i := 0; i < k; i++ { // 快指针 p1 先走 k 步
		p1 = p1.Next
	}
	p2 := dummy
	for p1 != nil { // 慢指针 p2 走了 length(list) - k 步
		p1 = p1.Next
		p2 = p2.Next
	}
	// 此时 p2 指向要删除节点上一个节点
	return p2
}
