package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// middleNode 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	l1 := head
	length := 0
	for l1 != nil {
		length++
		l1 = l1.Next
	}
	k := length / 2
	l1 = head
	for k > 0 {
		l1 = l1.Next
		k--
	}
	return l1
}

func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 假设链表长度为 k，当 fast 走到末尾时，slow 走到 k/2
		slow = slow.Next
	}
	return slow
}