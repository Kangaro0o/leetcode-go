package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// preorderTraversal 144. 二叉树的前序遍历
// 递归实现
func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	var preorder func(r *TreeNode)
	preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		res = append(res, r.Val)
		preorder(r.Left)
		preorder(r.Right)
	}
	preorder(root)
	return
}

// 迭代实现遍历树
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left // 因为是前序遍历，所以指向左子节点
		} else {
			// 弹出栈顶元素
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
