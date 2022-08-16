package sort

func quickSort(nums []int) {
	quickSort2(nums, 0, len(nums)-1)
}

func quickSort2(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(nums, l, r)
	quickSort2(nums, l, pivot-1)
	quickSort2(nums, pivot+1, r)
}

func partition(nums []int, l, r int) int {
	point := nums[r] // 选最后一个元素作为分区点
	i := l
	for j := l; j < r; j++ {
		if nums[j] < point {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i] // 处理最后一个元素，分区点会被放到"中间"位置
	return i
}
