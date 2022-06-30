package other

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	elem := nums[0]
	elemIdx := 0
	for i := 1; i < n; i++ {
		if elem == nums[i] { // 需要删除 i 位置上的元素
			for j := i; j < n-1; j++ {
				nums[j] = nums[j+1] // 删除了 nums[i] 位置的元素
			}
			n-- // 元素数量-1
			i = elemIdx
			// 有需要删除的元素时 elem 不需要更新
		} else { // 更新 elem
			elem = nums[i]
			elemIdx = i
		}
	}
	return n
}
