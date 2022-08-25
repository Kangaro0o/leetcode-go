package shunfeng

import "sort"

func sumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closestSum := sum(nums[0], nums[1], nums[2])
	if n == 3 {
		return closestSum
	}
	for i := 0; i < n-2; i++ {
		low, high := i+1, n-1
		for low < high {
			threeSum := sum(nums[i], nums[low], nums[high])
			if abs(threeSum-target) < abs(closestSum-target) {
				closestSum = threeSum
			}
			if threeSum > target {
				// 缩小范围
				high--
			} else if threeSum < target {
				low++
			} else {
				return threeSum
			}
		}
	}
	return closestSum
}

func sum(a, b, c int) int {
	return a + b + c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
