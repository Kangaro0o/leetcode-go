package bst

// convertBST 538. 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var preNode *TreeNode
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Right)
		if preNode == nil {
			preNode = node
		} else {
			node.Val = node.Val + preNode.Val
			preNode = node
		}
		inorder(node.Left)
		return
	}
	inorder(root)
	return root
}
