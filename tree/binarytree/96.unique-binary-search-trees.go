package binarytree

// 备忘录，记录计算结果，防止超时
var memo [][]int

func numTrees(n int) int {
	// 备忘录初始化
	memo = make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return build1(1, n)
}

func build1(lo, hi int) int {
	if lo > hi {
		return 1
	}
	if memo[lo][hi] != 0 { // 备忘录中存在，则直接使用
		return memo[lo][hi]
	}
	var res int
	for i := lo; i <= hi; i++ {
		leftNum := build1(lo, i-1)
		rightNum := build1(i+1, hi)
		res += leftNum * rightNum
	}
	memo[lo][hi] = res
	return res
}
