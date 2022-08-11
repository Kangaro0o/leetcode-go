package bst

// kthSmallest 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
		return
	}
	inorder(root)
	return res[k-1]
}
