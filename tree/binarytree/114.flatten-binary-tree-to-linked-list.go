package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// flatten 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	p := root
	res := preorder(root)
	for i := 1; i < len(res); i++ {
		next := &TreeNode{
			Val:  res[i],
			Left: nil,
		}
		p.Right = next
		p.Left = nil
		p = p.Right
	}
	return
}

func preorder(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var preorderHelper func(node *TreeNode)
	preorderHelper = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorderHelper(node.Left)
		preorderHelper(node.Right)
	}
	preorderHelper(root)
	return
}

// 在还没操作节点右子树前，不能破坏该节点的右子树指向。所以采用后序遍历。
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	flatten1(root.Left)
	flatten1(root.Right)
	left := root.Left
	right := root.Right
	// 将 root 的左子树和右子树拉平
	root.Left = nil
	root.Right = left
	// 将 root 的右子树接到左子树下方，然后将整个左子树作为右子树
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return
}
