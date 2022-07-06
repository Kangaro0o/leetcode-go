package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// maxDepth 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	// 类似树的后序遍历，等左子树和右子树节点树都计算出来后，返回最大值
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
