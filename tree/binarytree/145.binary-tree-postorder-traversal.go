package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// postorderTraversal 145. 二叉树的后序遍历
func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
		return
	}

	postorder(root)
	return
}
