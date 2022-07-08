package binarytree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect 116. 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	p := root
	queue := newNodeQueue()
	queue.offer(p)
	for !queue.isEmpty() {
		sz := queue.len()
		var q *Node // 用来完成指向 next 节点的操作
		for sz > 0 {
			p = queue.pop()
			if p.Left != nil {
				queue.offer(p.Left)
			}
			if p.Right != nil {
				queue.offer(p.Right)
			}
			if q == nil {
				q = p
			} else {
				q.Next = p
				q = p
			}
			sz--
		}
	}
	return root
}

type nodeQueue struct {
	data []*Node
}

func newNodeQueue() *nodeQueue {
	return &nodeQueue{}
}

func (q *nodeQueue) offer(elem *Node) {
	q.data = append(q.data, elem)
}

func (q *nodeQueue) pop() *Node {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *nodeQueue) len() int {
	return len(q.data)
}

func (q *nodeQueue) isEmpty() bool {
	return q.len() == 0
}
