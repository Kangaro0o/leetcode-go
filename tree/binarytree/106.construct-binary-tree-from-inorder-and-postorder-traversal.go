package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// buildTree1 106. 从中序与后序遍历序列构造二叉树
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree1(inorder[:i], postorder[:i])
	root.Right = buildTree1(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}
