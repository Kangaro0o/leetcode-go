package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res *TreeNode

// lowestCommonAncestor 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if isChild(root, p) && isChild(root, q) {
		res = root
	}
	lowestCommonAncestor(root.Left, p, q)
	lowestCommonAncestor(root.Right, p, q)
	return res
}

func isChild(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == p.Val {
		return true
	}
	return isChild(root.Left, p) || isChild(root.Right, p)
}
