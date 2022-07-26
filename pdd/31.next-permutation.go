package pdd

// nextPermutation 31. 下一个排列
// 从数组最后倒着查找，找到 nums[i] 比 nums[i+1] 小的时候，将两数互换，然后把 i+1, length-1数反转
func nextPermutation(nums []int) {
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		for j := length - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i] // 交换两数
				reverse(nums, i+1)
				return
			}
		}
	}
	reverse(nums, 0)
	return
}

func reverse(nums []int, start int) {
	cnt := (len(nums) - start) / 2
	for i := 0; i < cnt; i++ {
		nums[i+start], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i+start]
	}
}
