package other

import "sort"

// 16. 最接近的三数之和
func threeSumClosest1(nums []int, target int) int {
	n := len(nums)
	if n == 3 {
		return sum(nums[0], nums[1], nums[2])
	}
	s := sum(nums[0], nums[1], nums[2])
	delta := abs(s - target)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				s1 := sum(nums[i], nums[j], nums[k])
				if abs(s1-target) < delta {
					s = s1
					delta = abs(s1 - target)
				}
			}
		}
	}
	return s
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(a, b, c int) int {
	return a + b + c
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closestSum := sum(nums[0], nums[1], nums[2])
	if n == 3 {
		return closestSum
	}
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			threeSum := sum(nums[i], nums[left], nums[right])
			if abs(threeSum-target) < abs(closestSum-target) {
				closestSum = threeSum
			}
			if threeSum > target {
				// 缩小右侧搜索范围
				right--
			} else if threeSum < target {
				left++
			} else {
				return threeSum
			}
		}
	}
	return closestSum
}
