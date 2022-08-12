package bst

import "math"

// minDiffInBST 783. 二叉搜索树节点最小距离
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minVal := math.MaxInt
	var preNode *TreeNode
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if preNode != nil {
			minVal = min(node.Val-preNode.Val, minVal)
		}
		preNode = node
		inorder(node.Right)
		return
	}
	inorder(root)
	return minVal
}
