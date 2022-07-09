package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// countNodes 222. 完全二叉树的节点个数
func countNodes(root *TreeNode) (cnt int) {
	var count func(node *TreeNode)
	count = func(node *TreeNode) {
		if node == nil {
			return
		}
		count(node.Left)
		count(node.Right)
		cnt++
		return
	}
	count(root)
	return
}
