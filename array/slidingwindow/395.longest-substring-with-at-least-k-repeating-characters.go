package slidingwindow

import "strings"

// longestSubstring 395. 至少有 K 个重复字符的最长子串
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	need := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		need[s[i]]++
	}

	for key, val := range need {
		if val < k {
			var res int
			for _, subStr := range strings.Split(s, string(key)) {
				res = max(res, longestSubstring(subStr, k))
			}
			return res
		}
	}
	return len(s)
}
