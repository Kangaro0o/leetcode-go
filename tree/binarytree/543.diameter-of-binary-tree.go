package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// diameterOfBinaryTree 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepthHelper(root.Left)
	rightMax := maxDepthHelper(root.Right)
	return max(
		max(leftMax+rightMax, diameterOfBinaryTree(root.Left)),
		max(leftMax+rightMax, diameterOfBinaryTree(root.Right)),
	)
}

func maxDepthHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthHelper(root.Left), maxDepthHelper(root.Right)) + 1
}
