package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode

// 中序遍历，如果前一个节点比后一个值大，则返回 false
func isValidBST(root *TreeNode) bool {
	pre = nil
	return inorderBST(root)
}

func inorderBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := inorderBST(root.Left)
	if pre == nil {
		pre = root
	} else {
		preVal := pre.Val
		pre = root
		if preVal >= root.Val {
			return false
		}
	}
	right := inorderBST(root.Right)
	return left && right
}
