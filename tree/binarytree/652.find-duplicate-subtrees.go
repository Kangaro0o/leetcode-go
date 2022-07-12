package binarytree

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	nodeRes = nil
	treeMap = make(map[string]int)
	return find(root)
}

var treeMap map[string]int
var nodeRes []*TreeNode

func find(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var postorder func(node *TreeNode) string
	postorder = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		left := postorder(node.Left)
		right := postorder(node.Right)
		subTree := fmt.Sprintf("%s,%s,%s", left, right, fmt.Sprintf("%d", node.Val))
		if cnt, ok := treeMap[subTree]; ok && cnt == 1 {
			nodeRes = append(nodeRes, node)
		}
		treeMap[subTree]++
		return subTree
	}
	postorder(root)
	return nodeRes
}
