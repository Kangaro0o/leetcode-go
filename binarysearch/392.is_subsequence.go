package binarysearch

func isSubsequence(s string, t string) bool {
	ns, nt := len(s), len(t)
	if ns > nt {
		return false
	}
	if ns == 0 {
		return true
	}
	var p = 0 // 指向 s 中的元素
	for i := 0; i < nt; i++ {
		if s[p] == t[i] { // 如果能找到匹配的元素
			p++
		}
		if p == ns {
			return true
		}
	}
	return false
}
