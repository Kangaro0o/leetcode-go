package microsoft

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var pathSum int
	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		left := max(postOrder(node.Left), 0)
		right := max(postOrder(node.Right), 0)
		// 节点的最大路径和 取决于该节点的值与该节点的左右子节点的最大贡献值
		pathSum = max(node.Val+left+right, pathSum)
		// 因为路径没有岔路，返回给上一个节点的只能是路径，因此 max(left, right)
		return max(left, right) + node.Val
	}
	pathSum = math.MinInt
	postOrder(root)
	return pathSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
