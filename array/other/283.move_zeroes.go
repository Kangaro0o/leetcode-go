package other

import "fmt"

// 283. 移动零
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			t := nums[slow]
			nums[slow] = nums[fast]
			nums[fast] = t
			slow++
		}
		fast++
	}
	fmt.Println(nums)
}
