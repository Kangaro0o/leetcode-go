package bst

import "math"

func getMinimumDifference(root *TreeNode) int {
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
