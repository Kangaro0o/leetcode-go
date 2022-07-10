package narytree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) (res []int) {
	if root == nil {
		return nil
	}
	var postorderHelper func(node *Node)
	postorderHelper = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			postorderHelper(child)
		}
		res = append(res, node.Val)
		return
	}
	postorder(root)
	return
}
