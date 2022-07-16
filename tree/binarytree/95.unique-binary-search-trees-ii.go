package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// generateTrees 95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)
}

func build(lo, hi int) (res []*TreeNode) {
	if lo > hi {
		res = append(res, nil)
		return
	}
	for i := lo; i <= hi; i++ {
		leftTree := build(lo, i-1)
		rightTree := build(i+1, hi)
		// 固定左孩子，遍历右孩子
		for _, left := range leftTree {
			for _, right := range rightTree {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return
}
