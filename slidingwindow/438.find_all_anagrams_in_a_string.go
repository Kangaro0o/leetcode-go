package slidingwindow

// 该实现方式会出现超时问题
func findAnagrams1(s string, p string) []int {
	n := len(s)
	pLen := len(p)
	left, right := 0, pLen-1
	var res []int // 记录结果

	for right < n {
		right++
		s1 := s[left:right]
		if isAnagram(p, s1) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func isAnagram(a, b string) bool {
	set := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		set[a[i]]++
	}
	for j := 0; j < len(b); j++ {
		if _, ok := set[b[j]]; ok {
			set[b[j]]--
		}
	}
	for _, v := range set {
		if v != 0 {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	need := make(map[byte]int)
	pLen := len(p)
	for i := 0; i < pLen; i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	var res []int
	for right < len(s) {
		c := s[right]
		right++
		// 更新滑动窗口数据
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 	收缩窗口
		for (right - left) >= pLen {
			if valid == len(need) { // 有满足条件的子串
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}
