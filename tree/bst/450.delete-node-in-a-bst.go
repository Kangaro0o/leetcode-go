package bst

// deleteNode 450. 删除二叉搜索树中的节点
// 情况1：要删除的节点为叶子节点，则直接删除
// 情况2：要删除的节点有一个子节点，就让这个子节点接替自己的位置
// 情况3：要删除的节点有两个子节点，就选择左子树最大值或右子树最小值（选择其一即可）
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 情况1/2
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 情况3
		// 获取右子树最小节点
		minNode := getMinNode(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val) // 删除右子树最小节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMinNode(right *TreeNode) *TreeNode {
	for right.Left != nil {
		right = right.Left
	}
	return right
}
