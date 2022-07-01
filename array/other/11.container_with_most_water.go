package other

// maxArea 11. 盛最多水的容器
func maxArea(height []int) int {
	l := len(height)
	area := 0
	left, right := 0, l-1
	for left < right {
		t := (right - left) * minHeight(height[left], height[right])
		if t > area {
			area = t
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}

func minHeight(x, y int) int {
	if x > y {
		return y
	}
	return x
}
