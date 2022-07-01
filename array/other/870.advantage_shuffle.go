package other

import "sort"

// advantageCount 70. 优势洗牌
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	var res []int
	for j := 0; j < len(nums2); j++ {
		n := len(nums1)
		flag := false // 标识是否在 nums1 中找到比 nums2 中元素大的数
		for i := 0; i < n; i++ {
			if nums1[i] > nums2[j] { // nums1 中的数比 nums2 中当前数大时
				res = append(res, nums1[i])
				// 删除 nums1[i]
				nums1 = append(nums1[:i], nums1[i+1:]...)
				flag = true
				break
			}
		}
		if !flag {
			m, idx := min(nums1)
			nums1 = append(nums1[:idx], nums1[idx+1:]...)
			res = append(res, m)
		}
	}
	return res
}

func min(nums []int) (int, int) {
	m := nums[0]
	idx := 0
	for i, num := range nums {
		if num < m {
			m = num
			idx = i
		}
	}
	return m, idx
}
