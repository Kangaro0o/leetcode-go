package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// preorderTraversal 144. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var preorder func(r *TreeNode)
	preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		res = append(res, r.Val)
		preorder(r.Left)
		preorder(r.Right)
	}
	preorder(root)
	return
}
