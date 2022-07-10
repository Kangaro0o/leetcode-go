package binarytree

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 先序遍历序列化
	res := preorderSerialize(root)
	return strings.Join(res, ",")
}

func preorderSerialize(root *TreeNode) (res []string) {
	if root == nil {
		return nil
	}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			res = append(res, "#")
			return
		}
		res = append(res, fmt.Sprintf("%d", node.Val))
		preorder(node.Left)
		preorder(node.Right)
		return
	}
	preorder(root)
	return
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	s := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if s[0] == "#" {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		root := &TreeNode{
			Val: val,
		}
		s = s[1:] // left 和 right 递归共用此数组
		root.Left = build()
		root.Right = build()
		return root
	}
	return build()
}

func deserializeHelper(s []string) *TreeNode {
	if len(s) == 0 || s == nil {
		return nil
	}
	rootStrVal := s[0]
	if rootStrVal == "#" {
		s = s[1:]
		return nil
	}
	rootVal, _ := strconv.Atoi(rootStrVal)
	root := &TreeNode{Val: rootVal}
	s = s[1:]
	root.Left = deserializeHelper(s)
	root.Right = deserializeHelper(s)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
