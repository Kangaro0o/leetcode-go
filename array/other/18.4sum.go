package other

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n-3; i++ {
		// 三数之和
		threeSumRes := threeSumHelper(nums, i+1, target-nums[i])
		for _, ts := range threeSumRes {
			var t []int
			t = append(t, nums[i])
			t = append(t, ts...)
			res = append(res, t)
		}
		for i < n-3 && nums[i+1] == nums[i] {
			i++
		}
	}
	return res
}

// threeSumHelper 返回等于 target 的三数之和
// start nums数组开始的下标
// target 目标数
func threeSumHelper(nums []int, start, target int) [][]int {
	n := len(nums)
	var res [][]int
	for i := start; i < n-2; i++ {
		// 两数之和
		twoSumRes := twoSumHelper(nums, i+1, target-nums[i])
		for _, ts := range twoSumRes {
			var t []int
			t = append(t, nums[i])
			t = append(t, ts...)
			res = append(res, t)
		}
		for i < n-2 && nums[i+1] == nums[i] {
			i++
		}
	}
	return res
}

func twoSumHelper(nums []int, start, target int) [][]int {
	low, high := start, len(nums)-1
	var res [][]int
	for low < high {
		ts := nums[low] + nums[high]
		left, right := nums[low], nums[high]
		if ts > target {
			high--
		} else if ts < target {
			low++
		} else if ts == target {
			res = append(res, []int{nums[low], nums[high]})
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
