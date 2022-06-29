package binarysearch

import "sort"

// 俄罗斯套娃信封问题
// LIS 最长递增子序列
// 假设 envelopes[i] = [wi, hi]
func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(Envelop(envelopes))
	var height []int
	for _, e := range envelopes {
		height = append(height, e[1])
	}
	return getLengthOfLIS(height)
}

type Envelop [][]int

func (e Envelop) Len() int {
	return len(e)
}

// Less 实现 sort.Interface 接口的比较元素方法
func (e Envelop) Less(i, j int) bool {
	// 根据 [w, h] 中的 w 升序排序
	// 两种情况需要交换元素
	// 1. wi < wj
	// 2. wi = wj && hi > hj
	return e[i][0] < e[j][0] || (e[i][0] == e[j][0] && e[i][1] > e[j][1])
}

// Swap 实现 sort.Interface 接口的交换元素方法
func (e Envelop) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// getLengthOfLIS 查找最长递增子序列长度
func getLengthOfLIS(nums []int) int {
	// 使用二分查找排序找到对应位置，思路类似整理顺子牌
	n := len(nums)
	length := 0
	top := make([]int, n)
	for i := 0; i < n; i++ {
		poker := nums[i]
		left, right := 0, length
		// 二分查找插入位置
		for left < right {
			mid := left + (right-left)/2
			if top[mid] >= poker {
				// 如果 poker 比 top 中元素小，则说明 poker 比 top[x] 更有价值，后续存在比 poker 值大的数的概率会更大
				right = mid
			} else {
				left = mid + 1
			}
		}
		// 如果查找到最右侧，则说明 top 需要扩充
		if left == length {
			length++
		}
		top[left] = poker
	}
	// 最后的 top 不一定是最长子序列，但 length 一定是最长子序列长度
	return length
}
