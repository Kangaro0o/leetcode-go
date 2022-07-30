package alibaba

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// sortList 超时了
func sortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Val: -1}
	dummyHead.Next = head

	var cnt int
	for p := dummyHead.Next; p != nil; p = p.Next {
		cnt++
	}

	for i := 0; i < cnt; i++ {
		for p1 := dummyHead; p1 != nil && p1.Next != nil; p1 = p1.Next {
			first := p1.Next
			second := first.Next
			if second == nil {
				break
			}
			third := second.Next
			if first.Val > second.Val {
				first.Next = third
				second.Next = first
				p1.Next = second
			}
		}
	}
	return dummyHead.Next
}

// 使用归并排序的思想
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	var half *ListNode
	for fast != nil && fast.Next != nil {
		half = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	half.Next = nil
	left := mergeSort(head)
	right := mergeSort(slow)
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1}
	p := dummyHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	if left != nil {
		p.Next = left
	}
	if right != nil {
		p.Next = right
	}
	return dummyHead.Next
}
