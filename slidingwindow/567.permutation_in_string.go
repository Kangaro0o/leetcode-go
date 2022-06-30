package slidingwindow

// 567. 字符串的排列
// 给你两个字符串 s1 和s2，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	s1Len := len(s1)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok { // 如果 c 在 need 中，则新增
			window[c]++
			// 更新窗口数据
			if window[c] == need[c] {
				valid++
			}
		}

		// 缩小窗口操作
		for (right - left) >= s1Len {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok { // 如果 d 在 need 中，则减少
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
