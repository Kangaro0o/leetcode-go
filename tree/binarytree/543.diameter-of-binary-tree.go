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
// 递归+递归
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

var maxDiameter int

// 一次递归
func diameterOfBinaryTree1(root *TreeNode) int {
	maxDiameter = 0
	maxDepthHelper(root)
	return maxDiameter
}

func maxDepthHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepthHelper(root.Left)
	rightMax := maxDepthHelper(root.Right)
	maxDiameter = max(maxDiameter, leftMax+rightMax)
	return max(leftMax, rightMax) + 1
}
