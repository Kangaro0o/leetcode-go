package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// levelOrderBottom 107. 二叉树的层序遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := newQueue()
	queue.offer(root)
	for !queue.isEmpty() {
		sz := queue.len()
		var ans []int
		for sz > 0 {
			node := queue.pop()
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue.offer(node.Left)
			}
			if node.Right != nil {
				queue.offer(node.Right)
			}
			sz--
		}
		res = append(res, ans)
	}
	// 交换 res 数组顺序
	n := len(res)
	for i := 0; i < n/2; i++ {
		tmp := res[n-i-1]
		res[n-i-1] = res[i]
		res[i] = tmp
	}
	return res
}
