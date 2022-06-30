package slidingwindow

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0 // 记录结果
	n := len(s)
	for right < n {
		c := s[right]
		right++
		// 更新窗口数据
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
