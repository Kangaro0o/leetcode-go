package other

import "sort"

// threeSum 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return nil
	}
	var res [][]int
	for i := 0; i < n-2; i++ {
		num1 := nums[i]
		left, right := i+1, n-1
		for left < right {
			s := num1 + nums[left] + nums[right]
			if s > 0 {
				right--
			} else if s < 0 {
				left++
			} else {
				if exist := isExist(res, nums[i], nums[left], nums[right]); !exist {
					res = append(res, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--
			}
		}
	}
	return res
}

func isExist(res [][]int, a, b, c int) bool {
	for _, num := range res {
		if num[0] == a && num[1] == b && num[2] == c {
			return true
		}
	}
	return false
}
