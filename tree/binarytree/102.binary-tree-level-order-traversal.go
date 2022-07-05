package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// levelOrder 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int     // 保存每层节点
	queue := newQueue() // 使用队列保存
	queue.offer(root)
	for !queue.isEmpty() {
		var ans []int // 当前层元素值
		sz := queue.len()
		for sz > 0 {
			// 弹出栈顶元素
			top := queue.pop()
			ans = append(ans, top.Val)
			if top.Left != nil {
				queue.offer(top.Left)
			}
			if top.Right != nil {
				queue.offer(top.Right)
			}
			sz--
		}
		res = append(res, ans)
	}
	return res
}

type myQueue struct {
	data []*TreeNode
}

func newQueue() *myQueue {
	return &myQueue{}
}

func (q *myQueue) offer(elem *TreeNode) {
	q.data = append(q.data, elem)
}

func (q *myQueue) pop() *TreeNode {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *myQueue) len() int {
	return len(q.data)
}

func (q *myQueue) isEmpty() bool {
	return q.len() == 0
}
