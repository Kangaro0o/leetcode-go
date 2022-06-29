package binarysearch

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		} else if nums[mid] < target {
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1
}
