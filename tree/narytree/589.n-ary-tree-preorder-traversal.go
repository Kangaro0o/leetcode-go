package narytree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// preorder 589. N 叉树的前序遍历
func preorder(root *Node) []int {
	var res []int
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			traverse(child)
		}
		return
	}
	traverse(root)
	return res
}
