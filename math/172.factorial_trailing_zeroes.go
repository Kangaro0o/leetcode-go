package math

// 172. 阶乘后的零
func trailingZeroes1(n int) int {
	// 本质是求 5 的倍数
	var cnt int
	for n >= 5 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}

func trailingZeroes(n int) int {
	var cnt int
	for i := 5; i <= n; i += 5 {
		for j := i; j%5 == 0; j /= 5 {
			cnt++
		}
	}
	return cnt
}
