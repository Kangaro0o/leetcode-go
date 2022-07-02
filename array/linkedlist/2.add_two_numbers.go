// Package linkedlist 链表
package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 节点的数据结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	flag := false // 是否进位
	for l1 != nil || l2 != nil || flag {
		val := 0
		if flag {
			val += 1
			flag = false
		}
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val >= 10 {
			flag = true
			val = val % 10
		}
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 虚拟头节点
	p := dummy
	p1, p2 := l1, l2
	var carry int // 记录进位
	for p1 != nil || p2 != nil || carry > 0 {
		val := 0
		if carry > 0 {
			val += carry
		}
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}
		carry = val / 10
		val = val % 10
		// 记录该节点两数相加的值
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}
