package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// constructMaximumBinaryTree 654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	val, i := maxNum(nums)
	root := &TreeNode{Val: val}
	root.Left = constructMaximumBinaryTree(nums[:i])
	root.Right = constructMaximumBinaryTree(nums[i+1:])
	return root
}

// maxNum 返回数组中最大值和其下标
func maxNum(nums []int) (int, int) {
	m := nums[0]
	idx := 0
	for i, num := range nums {
		if num > m {
			m = num
			idx = i
		}
	}
	return m, idx
}
