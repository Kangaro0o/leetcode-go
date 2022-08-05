package bytedance

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var isOdd bool // 是否是偶数
	medianNumIdx := 0
	if (m+n)%2 == 0 { // 偶数
		isOdd = true
		medianNumIdx = (m+n)/2 - 1
	} else { // 奇数
		isOdd = false
		medianNumIdx = (m + n) / 2
	}
	count := 0
	i, j := 0, 0
	var medianNum []int
	for i < m && j < n {
		var median int
		if nums1[i] < nums2[j] {
			median = nums1[i]
			i++
		} else {
			median = nums2[j]
			j++
		}
		if count == medianNumIdx {
			medianNum = append(medianNum, median)
		}
		// 偶数需要额外处理
		if isOdd && count == medianNumIdx+1 {
			medianNum = append(medianNum, median)
		}
		count++
	}

	for i < m {
		median := nums1[i]
		if count == medianNumIdx {
			medianNum = append(medianNum, median)
		}
		// 偶数需要额外处理
		if isOdd && count == medianNumIdx+1 {
			medianNum = append(medianNum, median)
		}
		count++
		i++
	}

	for j < n {
		median := nums2[j]
		if count == medianNumIdx {
			medianNum = append(medianNum, median)
		}
		// 偶数需要额外处理
		if isOdd && count == medianNumIdx+1 {
			medianNum = append(medianNum, median)
		}
		count++
		j++
	}

	if isOdd {
		return float64(medianNum[0]+medianNum[1]) / 2
	}
	return float64(medianNum[0])
}
