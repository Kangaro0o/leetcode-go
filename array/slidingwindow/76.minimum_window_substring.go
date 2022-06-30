package slidingwindow

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	res := "" // 记录结果
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok { // 如果 c 在 need 中存在，则新增
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 缩小窗口
		for valid == len(need) {
			d := s[left]
			if _, ok := need[d]; ok { // 如果 d 在 need 中存在，则整体向右滑动
				if window[d] == need[d] {
					valid--
				}
				window[d]--
				// 这里要确定 left 对应的值在 need 中存在
				p := s[left:right]
				if res == "" || len(p) < len(res) {
					res = p
				}
			}
			// 确定 left 滑动到哪
			for i := left + 1; i < right; i++ {
				if _, ok := need[s[i]]; ok {
					left = i
					break
				}
			}
		}
	}
	return res
}
