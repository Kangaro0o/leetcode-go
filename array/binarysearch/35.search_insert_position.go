package binarysearch

// 这题本质是在 nums 中查找匹配 target 的左边界
func searchInsert(nums []int, target int) int {
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
	if left < len(nums) && nums[left] >= target {
		return left
	}
	return left + 1
}
