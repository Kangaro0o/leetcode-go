package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// findMode 501. 二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	var res []int
	cntMap := make(map[int]int)
	maxCnt := 0
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		cntMap[node.Val]++
		cnt := cntMap[node.Val]
		if cnt > maxCnt {
			maxCnt = cnt
			// 删除 res 中之前保存的众数
			res = nil
			// 添加到结果中
			res = append(res, node.Val)
		} else if cnt == maxCnt {
			res = append(res, node.Val)
		}
		preorder(node.Left)
		preorder(node.Right)
		return
	}
	preorder(root)
	return res
}
