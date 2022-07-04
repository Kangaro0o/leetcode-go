package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// detectCycle 142. 环形链表 II
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		// 排除 fast 遍历到末尾结束循环的可能性
		return nil
	}
	// 说明链表有环
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
