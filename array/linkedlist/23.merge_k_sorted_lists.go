package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// mergeKLists 23. 合并K个升序链表
// 循环合并
func mergeKLists1(lists []*ListNode) *ListNode {
	var p *ListNode
	for i := 0; i < len(lists); i++ {
		p = mergeTwoList(p, lists[i])
	}
	return p
}

// mergeKLists 23. 合并K个升序链表
//
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	p1 := merge(lists, 0, mid)
	p2 := merge(lists, mid+1, right)
	return mergeTwoList(p1, p2)
}

// mergeTwoList 合并两个链表
func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		val := 0
		if p1.Val < p2.Val {
			val = p1.Val
			p1 = p1.Next
		} else {
			val = p2.Val
			p2 = p2.Next
		}
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
