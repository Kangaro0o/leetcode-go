package other

// trap 42. 接雨水
// 对于任意一个位置的柱子i，能够装的水为 water[i] = min(max(height[0...i]), max(height(i...end))) - height[i]
// 在于如何能快速计算出某个位置左侧所有柱子的最大高度和右侧所有柱子的最大高度
func trap(height []int) int {
	n := len(height)
	water := 0
	for i := 1; i < n-1; i++ {
		m1 := maxHeightHelper(height[:i])
		m2 := maxHeightHelper(height[i+1:])
		t := minHeight(m1, m2) - height[i] // 能装的水
		if t > 0 {
			water += t
		}
	}
	return water
}

func maxHeightHelper(height []int) int {
	max := 0
	for _, h := range height {
		if h > max {
			max = h
		}
	}
	return max
}
