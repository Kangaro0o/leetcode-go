package binarysearch

func nthUglyNumber(n int, a int, b int, c int) int {
	var res = 0
	for i := 0; i < n; i++ {
		var num = res + 1
		for {
			if num%a == 0 || num%b == 0 || num%c == 0 {
				res = num
				break
			}
			num++
		}
	}
	return res
}
