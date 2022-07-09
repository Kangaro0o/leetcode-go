package binarytree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// tree2str 606. 根据二叉树创建字符串
func tree2str(root *TreeNode) (res string) {
	if root == nil {
		return
	}
	var preo func(node *TreeNode)
	preo = func(node *TreeNode) {

		res += fmt.Sprintf("%d", node.Val)

		if node.Left != nil {
			res += "("
			preo(node.Left)
			res += ")"
		}
		// 这步很关键，处理了root = [1,2,3,null,4]这类case
		if node.Left == nil && node.Right != nil {
			res += "()"
		}
		if node.Right != nil {
			res += "("
			preo(node.Right)
			res += ")"
		}

		return
	}
	preo(root)
	return
}
