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
