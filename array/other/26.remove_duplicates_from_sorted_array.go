package other

// 26. 删除有序数组中的重复项
func removeDuplicates1(nums []int) int {
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

// 使用快慢指针
// 0 0 1 1 2
func removeDuplicates(nums []int) int {
	// 使用快慢指针，慢指针走在后面，快指针走在前面探路，找到一个不重复的元素就赋值给慢指针，并让慢指针前进一步
	// 这样就保证了 nums[0...slow]都是无重复元素，当快指针遍历完整个数组后，程序结束
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
