package slidingwindow

func LengthOfLongestSubstringTwoDistinct(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	var res int
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for len(window) > 2 {
			d := s[left]
			left++
			window[d]--
			if window[d] == 0 {
				delete(window, d)
			}
		}
		res = max(right-left, res)
	}
	return res
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
