package bst

// findMode 501. 二叉搜索树中的众数
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	m := make(map[int]int)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 优化点：在此处判断出现频率最高的元素并记录
		if _, ok := m[node.Val]; ok {
			m[node.Val]++
		} else {
			m[node.Val] = 1
		}
		preorder(node.Left)
		preorder(node.Right)
		return
	}
	preorder(root)
	return getMode(m)
}

func getMode(m map[int]int) []int {
	var maxVal int
	for _, v := range m {
		if v > maxVal {
			maxVal = v
		}
	}
	var res []int
	for k, v := range m {
		if v == maxVal {
			res = append(res, k)
		}
	}
	return res
}
