package slidingwindow

// maxSubArray 53. 最大子数组和
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	left, right := 0, 0
	res := 0
	max := nums[0]
	for right < n {
		// 扩大窗口
		c := nums[right]
		right++
		res += c
		if res > max {
			max = res
		}
		// 缩小窗口
		for res < 0 {
			d := nums[left]
			left++
			res = res - d
		}
	}
	return max
}
