package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// getIntersectionNode1 160. 相交链表
// p1 p2 轮流走 headA headB 直到相同
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	l1, l2 := headA, headB
	for l1 != nil {
		lenA++
		l1 = l1.Next
	}
	for l2 != nil {
		lenB++
		l2 = l2.Next
	}
	l1, l2 = headA, headB
	if lenA > lenB {
		delta := lenA - lenB
		for delta > 0 {
			l1 = l1.Next
			delta--
		}
	} else {
		delta := lenB - lenA
		for delta > 0 {
			l2 = l2.Next
			delta--
		}
	}
	for l1 != l2 {
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil || l2 == nil {
		return nil
	}
	return l1
}
