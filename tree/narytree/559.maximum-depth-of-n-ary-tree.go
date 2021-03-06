// Package narytree N 叉树相关算法
package narytree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queue := NewQueue()
	queue.offer(root)
	var maxDepthNum int
	for !queue.isEmpty() {
		sz := queue.size()
		maxDepthNum++
		for sz > 0 {
			root = queue.pop()
			for _, child := range root.Children {
				queue.offer(child)
			}
			sz--
		}
	}
	return maxDepthNum
}

// 递归遍历计数
func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}
	var maxDepthNum int
	for _, child := range root.Children {
		childDepth := maxDepth1(child)
		maxDepthNum = max(childDepth, maxDepthNum)
	}
	return maxDepthNum + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func NewQueue() *MyQueue {
	return &MyQueue{}
}

type MyQueue struct {
	data []*Node
}

func (q *MyQueue) offer(elem *Node) {
	q.data = append(q.data, elem)
}

func (q *MyQueue) pop() *Node {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *MyQueue) size() int {
	return len(q.data)
}

func (q *MyQueue) isEmpty() bool {
	return q.size() == 0
}
