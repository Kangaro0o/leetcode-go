package bytedance

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	var ans int
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
