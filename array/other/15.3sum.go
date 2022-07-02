package other

import (
	"sort"
)

// threeSum 15. 三数之和
func threeSum1(nums []int) [][]int {
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

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		twoSumRes := twoSum(nums, i+1, 0-nums[i])
		for _, twoRes := range twoSumRes {
			var t []int
			t = append(t, nums[i])
			t = append(t, twoRes...)
			if t != nil {
				res = append(res, t)
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(nums []int, start, target int) [][]int {
	low, high := start, len(nums)-1
	var res [][]int
	for low < high {
		sum := nums[low] + nums[high]
		left, right := nums[low], nums[high]
		if sum > target {
			for low < high && nums[high] == right {
				high--
			}
		} else if sum < target {
			for low < high && nums[low] == left {
				low++
			}
		} else if sum == target {
			res = append(res, []int{nums[low], nums[high]})
			// 去重
			for low < high && nums[low] == left {
				low++
			}
			for low < high && nums[high] == right {
				high--
			}
		}
	}
	return res
}
