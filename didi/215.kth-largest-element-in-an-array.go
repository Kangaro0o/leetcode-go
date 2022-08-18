package didi

func findKthLargest(nums []int, k int) int {
	nums = mergeSort(nums)
	return nums[k-1]
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	var res []int
	for i < m && j < n {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	for i < m {
		res = append(res, left[i])
		i++
	}

	for j < n {
		res = append(res, right[j])
		j++
	}
	return res
}
