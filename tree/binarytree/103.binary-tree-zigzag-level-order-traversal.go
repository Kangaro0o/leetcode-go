package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// zigzagLevelOrder 103. 二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := newQueue()
	queue.offer(root)
	leftToRight := true
	for !queue.isEmpty() {
		sz := queue.len()
		var ans []int
		for sz > 0 {
			// 弹出队头元素
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
		var arr []int
		if leftToRight { // 从左向右
			for i := 0; i < len(ans); i++ {
				arr = append(arr, ans[i])
			}
		} else {
			for i := len(ans) - 1; i >= 0; i-- {
				arr = append(arr, ans[i])
			}
		}
		leftToRight = !leftToRight
		res = append(res, arr)
	}
	return res
}
