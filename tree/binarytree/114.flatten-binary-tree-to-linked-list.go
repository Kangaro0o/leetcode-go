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
