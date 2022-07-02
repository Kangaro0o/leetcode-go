// Package binarysearch 二分搜索
package binarysearch

func searchRange(nums []int, target int) []int {
	return []int{getLeftBound(nums, target), getRightBound(nums, target)}
}

func getLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		} else if nums[mid] < target {
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		} else if nums[mid] == target {
			right = mid
		}
	}
	// 检查边界
	// nums[left] != target 是防止 left == right 时，对应值不等于 target
	// left > len(nums) 是检查 nums 是否为空
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left // 跳出循环时, left == right 且对应的值为target,即 left或者right 为 target 的左边界
}

func getRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		// 有一种特殊情况：当 left == right-1 && nums[mid] = target 时
		// 此时 left 不会有任何变化，会进入死循环，因此 mid 需要 + 1
		mid := left + (right-left)/2 + 1
		if nums[mid] > target {
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		} else if nums[mid] < target {
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		} else if nums[mid] == target {
			// 不能直接返回，因为要找出右边界，所以left需要右移，但不能left=mid+1,因为这样的话可能就不再满足循环条件
			// 我们要找到 left == right 且对应值为 target 的情况
			left = mid
		}
	}
	// 检查边界
	// nums[right] != target 是防止 left == right 时，对应值不等于 target
	// right < 0 是检查 nums 是否为空
	if right < 0 || nums[right] != target {
		return -1
	}
	return right // 跳出循环时，left==right 且对应值为target值，即 left或right为target的右边界
}
