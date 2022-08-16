package bst

// maxSumBST 1373. 二叉搜索子树的最大键值和
func maxSumBST(root *TreeNode) int {
	res = 0
	return helper(root)
}

var res int

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	helper(root.Left)
	helper(root.Right)
	if ok, maxSum := isValidBST1(root); ok {
		res = max(maxSum, res)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isValidBST1(root *TreeNode) (valid bool, maxSum int) {
	if root == nil {
		return true, 0
	}
	var preNode *TreeNode
	var preorder func(node *TreeNode) bool
	preorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		l := preorder(node.Left)

		if preNode != nil && preNode.Val >= node.Val {
			return false
		}
		maxSum += node.Val
		preNode = node

		r := preorder(node.Right)
		return l && r
	}
	return preorder(root), maxSum
}
