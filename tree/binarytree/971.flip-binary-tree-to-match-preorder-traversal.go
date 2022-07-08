package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// flipMatchVoyage 971. 翻转二叉树以匹配先序遍历
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	if root == nil {
		return nil
	}
	var ans []int
	var pos int // 记录遍历的节点对应 voyage 的位置
	var stack []*TreeNode
	diffMap := make(map[int]bool)
	for root != nil || len(stack) > 0 {
		if len(stack) == 0 && root.Val != voyage[pos] {
			return []int{-1}
		}
		if root != nil {
			if root.Val == voyage[pos] {
				// 当前节点和先序遍历节点相同，则跳到下一节点
				stack = append(stack, root)
				root = root.Left
				pos++
			} else {
				// 记录不相同的元素
				if diffMap[root.Val] { // 此时不管如何交换左右节点都不可能相等
					return []int{-1}
				} else {
					diffMap[root.Val] = true
				}
				// 按 root 进行翻转
				root = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				pos--
				// 记录翻转节点
				ans = append(ans, root.Val)
				tmp := root.Left
				root.Left = root.Right
				root.Right = tmp
			}
		} else {
			// 弹出栈顶元素
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return ans
}
