package microsoft

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
	l, r := len(left), len(right)
	i, j := 0, 0
	var res []int
	for i < l && j < r {
		if left[i] > right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	for i < l {
		res = append(res, left[i])
		i++
	}
	for j < r {
		res = append(res, right[j])
		j++
	}
	return res
}
