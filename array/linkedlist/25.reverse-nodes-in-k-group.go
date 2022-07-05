package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// reverseKGroup 25. K 个一组翻转链表
// 1->2->3->4->5->6->7 k=3
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	length := 0 // 记录链表节点数
	for head != nil {
		length++
		head = head.Next
	}
	// 计算需要反转的子链表个数
	prev := dummy
	curr := dummy.Next
	for i := 0; i < length/k; i++ {
		for i := 0; i < k-1; i++ {
			next := curr.Next   // 下一个节点
			second := next.Next // 下两个节点
			curr.Next = second
			next.Next = prev.Next // 指向链表头节点
			prev.Next = next      // 更新链表头节点
		}
		prev = curr // 重新定位新的 prev 和 curr 指针
		curr = prev.Next
	}
	return dummy.Next
}

// 递归解法
func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 翻转前 k 个元素
	newHeadNode := reverse(a, b)
	// 翻转后 a 成为前 k 个元素中最后一个节点
	//递归翻转后续链表并链接起来
	a.Next = reverseKGroup1(b, k)
	return newHeadNode
}

// 翻转[a,b) 区间的元素，左闭右开
func reverse(a, b *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = a
	curr := a
	prev := dummy
	for curr.Next != b {
		next := curr.Next
		second := next.Next
		next.Next = prev.Next
		curr.Next = second
		prev.Next = next // 重新设置头节点
	}
	return prev.Next
}
